package db

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"reqyt.run/proxy/config"
)

var DB *sqlx.DB

func InitDB() {
	db, err := sqlx.Open("sqlite3", config.Config.DatabasePath)
	if err != nil {
		log.Fatal("Could not open DB", err)
		return
	}
	DB = db

	log.Info("DB OK")
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Error("Error calling DB.Close()", err)
	} else {
		log.Info("DB Closed")
	}
}