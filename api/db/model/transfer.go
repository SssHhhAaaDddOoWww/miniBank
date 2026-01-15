package model

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	FromAccountID uint    `gorm:"not null"`
	ToAccountID   uint    `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
}
