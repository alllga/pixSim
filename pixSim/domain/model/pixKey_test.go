package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/alllga/pixSim/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "j@j.com"
	pixKey, err := model.NewPixKey(account, key, kind)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, err = model.NewPixKey(account, key, kind)
	require.Nil(t, err)

	_, err = model.NewPixKey(account, key, "nome")
	require.NotNil(t, err)
}