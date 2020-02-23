package db

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"reqyt.run/core_api/config"
)

var DB *sqlx.DB

type LogWriter struct{}

func (lw *LogWriter) Write(p []byte) (n int, err error) {
	log.Info("SqlBoiler: " + string(p))
	return len(p), nil
}

func InitDB() {
	db, err := sqlx.Open("sqlite3", config.Config.DatabasePath)
	if err != nil {
		log.Fatal("Could not open DB", err)
		return
	}
	DB = db
	boil.SetDB(db)
	if config.Config.LogQueries {
		boil.DebugMode = true
		boil.DebugWriter = &LogWriter{}
	}
	// initial schema
	if _, err := DB.Exec(Schema); err != nil {
		log.Warn(err)
	}

	// initial data
	if _, err := DB.Exec(Data); err != nil {
		log.Warn(err)
	}
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