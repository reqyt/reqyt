package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"reqyt.run/proxy/config"
	"reqyt.run/proxy/db"
	"strings"
	"syscall"
	"time"
)

func setupLogging() {
	log.SetReportCaller(config.Config.LogCaller)
	level, err := log.ParseLevel(config.Config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)
}


type Function struct {
	ID          int64
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
	Slug        string
	Description string
	Name        string
	Cost        int64
	Endpoint    string
	WalletID    int64 `db:"wallet_id"`
}

type Wallet struct {
	ID          int64
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
	Key         string
	Credits     int64
	Endpoint    string
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	requestedFunc := strings.ReplaceAll(r.RequestURI, "/", "")
	funcRow := db.DB.QueryRowx("SELECT * FROM functions WHERE slug=?", requestedFunc)
	function := &Function{}
	err := funcRow.StructScan(function)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	walletID := 2
	walletRow := db.DB.QueryRowx("SELECT * FROM wallets WHERE id=?", walletID)
	wallet := &Wallet{}
	err = walletRow.StructScan(wallet)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// credits are unsigned integer in sqlite, i would expect it to fail <0 but it doesn't.
	if wallet.Credits < function.Cost {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	tx, err := db.DB.Begin()
	tx.Exec("UPDATE wallets SET credits = credits-? WHERE wallets.id = ?", function.Cost, wallet.ID)
	tx.Exec("UPDATE wallets SET credits = credits+? WHERE wallets.id = ?", function.Cost, function.WalletID)
	tx.Exec(`
		INSERT INTO function_calls(created_at, updated_at, function_id, wallet_id, credits)
		VALUES(time('now'), time('now'), ?,?,?);
		`, function.ID, wallet.ID, function.Cost)
	err = tx.Commit()

	log.Info("Forwarding request to: ", function.Endpoint)
	req, err := http.NewRequest("POST", function.Endpoint, r.Body)
	req.Header.Set("content-type", r.Header.Get("content-type"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		log.Warn(err)
		return
	}

	w.Header().Set("content-type", resp.Header.Get("content-type"))
	io.Copy(w, resp.Body)
	log.Info("Completed!")

}
func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}
	setupLogging()
	db.InitDB()
	defer db.CloseDB()
	server := &http.Server{Addr: config.Config.ApiBindAddress, Handler: http.HandlerFunc(ProxyHandler)}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		s := <-c
		log.Warn("Got signal: ", s)
		server.Close()
	}()
	log.Info("Starting API server on ", config.Config.ApiBindAddress)
	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
		log.Warn(err)
	}
}
