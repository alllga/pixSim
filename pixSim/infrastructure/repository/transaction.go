package repository

import (
	"fmt"

	"github.com/alllga/pixSim/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

// type TransactionRepositoryInterface interface {
// 	Register(transaction *Transaction) error
// 	Save(transaction *Transaction) error
// 	Find(id string) (*Transaction, error)
// }


func (trans *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := trans.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (trans *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := trans.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (trans *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var foundTransaction model.Transaction
	trans.Db.Preload("AccountFrom.Bank").First(&foundTransaction, "id = ?", id)

	if foundTransaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}
	return &foundTransaction, nil

}