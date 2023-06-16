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
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Ammount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyIdTo        string   `gorm:"column:pix_key_id_to;type:uuid;" valid:"notnull"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"-"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
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


func NewTransaction(accountFrom *Account, ammount float64, pixKeyTo *PixKey, description string, id string) (*Transaction, error) {
	transAc := Transaction{
		AccountFrom:   accountFrom,
		AccountFromID: accountFrom.ID,
		Ammount:        ammount,
		PixKeyTo:      pixKeyTo,
		PixKeyIdTo:    pixKeyTo.ID,
		Status:        TransactionPending,
		Description:   description,
	}

	if id == "" {
		transAc.ID = uuid.NewV4().String()
	} else {
		transAc.ID = id
	}
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