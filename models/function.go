package models

import "github.com/jinzhu/gorm"

type Source struct {
	gorm.Model
	URL      string
	Function uint
}

// Function bla
type Function struct {
	gorm.Model
	Slug   string
	Name   string
	Cost   int64
	Source Source  `gorm:"foreignkey:Function"`
	Tokens []Token `gorm:"many2many:rel_token_function;"`
}
