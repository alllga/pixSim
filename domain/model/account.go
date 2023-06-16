package model

import (
	"time"

	valid "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)


type Account struct {
	Base				`valid:"required"`
	OwnerName		string	`json:"owner_name" valid:"notonull"`
	Bank				*Bank	  `valid:"-"`
	Number			string	`json:"number" valid:"notnull"`
	PixKeys			[]*PixKey `valid:"-"`
}

func (account *Account) isValid() error {
	_, err := valid.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}


func newAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Bank: bank,
		Number: number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil 
}