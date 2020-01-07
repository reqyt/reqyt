package models

import "github.com/jinzhu/gorm"

// User bla
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Credits   int64
}
