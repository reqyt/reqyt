package main

import (
	"marktwolk/core/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, _ := gorm.Open("sqlite3", "./gorm.db")
	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.Function{}, &models.Source{}, &models.Token{})
	p1 := models.Function{}
	db.Create(&p1)

	p2 := models.Source{Function: p1.ID}
	db.Create(&p2)

}
