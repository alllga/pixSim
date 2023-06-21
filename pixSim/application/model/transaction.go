package model

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountId" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"-"`
	Error        string  `json:"error"`
}

func (trans *Transaction) isValid() error {
	valid := validator.New()
	err := valid.Struct(trans)
	if err != nil {
		fmt.Println("Error during Transaction validation:", err)
		return err
	}
	return nil
}

func (trans *Transaction) ParseJson(data []byte) error {
	err := json.Unmarshal(data, trans)
	if err != nil {
		return err
	}

	err = trans.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (trans *Transaction) ToJson() ([]byte, error) {
	err := trans.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(trans)
	if err != nil {
		return nil, nil
	}

	return result, nil
}

func NewTransaction() *Transaction {
	return &Transaction{}
}