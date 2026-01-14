package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	balance float64 `gorm:"not null; @default:0"`
}
