package main

import (
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"reqyt.run/core_api/config"
	"reqyt.run/core_api/db"
	"reqyt.run/core_api/rest_api"
	"syscall"
)

func setupLogging() {
	log.SetReportCaller(config.Config.LogCaller)
	level, err := log.ParseLevel(config.Config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)
}


func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}
	setupLogging()
	db.InitDB()
	defer db.CloseDB()
	server := &http.Server{Addr: config.Config.ApiBindAddress, Handler: rest_api.CreateRouter()}
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
