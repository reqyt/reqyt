package models_test

import (
	"fmt"
	"marktwolk/core/models"

	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestSum(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "./gorm.db")
	defer db.Close()
	db.AutoMigrate(&models.Person{})
	p1 := models.Person{FirstName: "John", LastName: "Doe"}
	p2 := models.Person{FirstName: "Jane", LastName: "Smith"}

	db.Create(&p1)
	var p3 models.Person // identify a Person type for us to store the results in
	db.First(&p3)        // Find the first record in the Database and store it in p3

	fmt.Println(p1.FirstName)
	fmt.Println(p2.LastName)
	fmt.Println(p3.LastName)
}
