package model

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	FromAccountID uint  `gorm:"not null"`
	ToAccountID   uint  `gorm:"not null"`
	Amount        int64 `gorm:"not null"`
}
