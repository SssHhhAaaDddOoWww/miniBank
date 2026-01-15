package services

import (
	"errors"

	"github.com/SssHhhAaaDddOoWww/miniBank/internal/db"
	"github.com/SssHhhAaaDddOoWww/miniBank/internal/db/model"
	"gorm.io/gorm"
)

func Deposit(amount float64, accountID uint) error {
	if amount <= 0 {
		return errors.New("amount must be positive !")

	}
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var acc model.Account
		err := tx.First(&acc, accountID).Error
		if err != nil {
			return err
		}
		acc.Balance += amount
		if err := tx.Save(&acc).Error; err != nil {
			return err
		}
		entry := model.LedgerEntry{
			AccountID: accountID,
			Amount:    float64(amount),
		}
		if err = tx.Create(&entry).Error; err != nil {
			return err
		}
		return nil
	})
}

func Withdraw(amount float64, accountID uint) error {
	if amount <= 0 {
		return errors.New("amount must be positive !")
	}
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var acc model.Account
		err := tx.First(&acc, accountID).Error
		if err != nil {
			return err
		}
		acc.Balance -= amount
		if err := tx.Save(&acc).Error; err != nil {
			return err
		}
		entry := model.LedgerEntry{
			AccountID: accountID,
			Amount:    -amount,
		}
		if err = tx.Create(&entry).Error; err != nil {
			return err
		}
		return nil
	})

}
func Transfer(fromID, toID uint, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	if fromID == toID {
		return errors.New("cannot transfer to same account")
	}

	return db.DB.Transaction(func(tx *gorm.DB) error {
		var from, to model.Account

		if err := tx.First(&from, fromID).Error; err != nil {
			return err
		}

		if err := tx.First(&to, toID).Error; err != nil {
			return err
		}

		if from.Balance < (amount) {
			return errors.New("insufficient balance")
		}

		transfer := model.Transfer{
			FromAccountID: from.ID,
			ToAccountID:   to.ID,
			Amount:        amount,
		}

		if err := tx.Create(&transfer).Error; err != nil {
			return err
		}

		from.Balance -= amount
		to.Balance += amount

		if err := tx.Save(&from).Error; err != nil {
			return err
		}

		if err := tx.Save(&to).Error; err != nil {
			return err
		}

		debit := model.LedgerEntry{
			AccountID:  from.ID,
			Amount:     -amount,
			TransferID: &transfer.ID,
		}

		credit := model.LedgerEntry{
			AccountID:  to.ID,
			Amount:     float64(amount),
			TransferID: &transfer.ID,
		}

		if err := tx.Create(&debit).Error; err != nil {
			return err
		}

		if err := tx.Create(&credit).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetBalance(accountID uint) (float64, error) {
	var acc model.Account

	if err := db.DB.First(&acc, accountID).Error; err != nil {
		return 0, err
	}

	return acc.Balance, nil
}

func GetLedger(accountID uint) ([]model.LedgerEntry, error) {
	var entries []model.LedgerEntry
	if err := db.DB.Where("account_id = ?", accountID).
		Order("created_at desc").
		Find(&entries).Error; err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, errors.New("no ledger entries found")
	}

	return entries, nil
}
