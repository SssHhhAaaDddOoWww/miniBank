package model

import "gorm.io/gorm"

type LedgerEntry struct {
	gorm.Model

	AccountID uint  `gorm:"not null"`
	Amount    int64 `gorm:"not null"`

	TransactionID uint
	TransferID    *uint
}
