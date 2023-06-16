package model

import (
	"errors"
	"time"

	valid "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(account *Account) (*Account, error)
}

type PixKey struct {
	Base				`valid:"required"`
	Kind				string	`json:"kind" valid:"notnull"`
	Key					string	`json:"key" valid:"notnull"`
	AccountId		string	`json:"account_id" valid:"notnull"`
	Account			*Account	`valid:"-"`
	Status			string	`json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := valid.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid type of status")
	}

	if err != nil {
		return err
	}

	return nil
}


func NewPixKey(account *Account, key string, kind string) (*PixKey, error) {

	pixKey := PixKey{
		Kind: kind,
		Key: key,
		Account: account,
		AccountId: account.ID,
		Status: "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	pixKey.UpdatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil 
}