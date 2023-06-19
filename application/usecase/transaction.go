package usecase

import (
	"errors"
	"log"

	"github.com/alllga/pixSim/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository 				model.PixKeyRepositoryInterface
}

func (transCase *TransactionUseCase) Register(accountId string, ammount float64, pixKeyTo string, pixKeyKindTo, description string) (*model.Transaction, error) {

	account, err := transCase.PixRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := transCase.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, ammount, pixKey, description, "")
	if err != nil {
		return nil, err
	}

	transCase.TransactionRepository.Save(transaction)
	if transaction.ID == "" {
		return nil, errors.New("unable to start new transaction at the moment")
	}

	return transaction, nil

}

func (transCase *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := transCase.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = transCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (transCase *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := transCase.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err = transCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (transCase *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := transCase.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = transCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

