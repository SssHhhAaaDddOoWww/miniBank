package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Type      string `gorm:"not null"`
	Amount    int64  `gorm:"not null"`
	AccountID uint   `gorm:"not null"`
}
