package models

import "github.com/jinzhu/gorm"

// Token bla
type Token struct {
	gorm.Model
	Hash      string
	User      uint
	Credits   int64
	Functions []Function `gorm:"many2many:rel_token_function;"`
}
