package model

import (
	"errors"
	"time"

	valid "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)


const (
	TransactionPending string = "pending"
	TransactionCompleted string = "completed"
	TransactionError string = "error"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base							`valid:"required"`
	AccountFrom				*Account `valid:"-"`
	Ammount						float64 `json:"ammount" valid:"required"`
	PixKeyTo					*PixKey `valid:"-"`
	Status						string `json:"status" valid:"notnull"`
	Description				string `json:"description" valid:"notnull"`
	CancelDescription	string `json:"cancel_description" valid:"-"`
}

func (transAc *Transaction) isValid() error {
	_, err := valid.ValidateStruct(transAc)


	if transAc.Ammount <= 0 {
		return errors.New("the ammount must be greater than 0")
	}

	if transAc.Status != TransactionPending && transAc.Status != TransactionConfirmed && transAc.Status != TransactionError {
		return errors.New("invalid transation status")
	}

	if transAc.PixKeyTo.AccountId == transAc.AccountFrom.ID {
		return errors.New("the source and destination cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}


func newTransaction(accountFrom *Account, ammount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transAc := Transaction{
		AccountFrom: accountFrom,
		Ammount: ammount,
		PixKeyTo: pixKeyTo,
		Description: description,
		Status: TransactionPending,
	}

	transAc.ID = uuid.NewV4().String()
	transAc.CreatedAt = time.Now()
	transAc.UpdatedAt = time.Now()

	err := transAc.isValid()
	if err != nil {
		return nil, err
	}

	return &transAc, nil 
}

func (transAc *Transaction) Complete()  error {
	transAc.Status = TransactionCompleted
	transAc.UpdatedAt = time.Now()
	err := transAc.isValid()
	return err
}

func (transAc *Transaction) Cancel(description string)  error {
	transAc.Status = TransactionError
	transAc.UpdatedAt = time.Now()
	transAc.Description = description
	err := transAc.isValid()
	return err
}

func (transAc *Transaction) Confirm()  error {
	transAc.Status = TransactionConfirmed
	transAc.UpdatedAt = time.Now()
	err := transAc.isValid()
	return err
}