package model

import (
	"time"

	valid "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base `valid:"required"`
	Code string `json:"code" valid:"notnull"`
	Name string `json:"name" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:BankID" valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := valid.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	bank.UpdatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil 
}