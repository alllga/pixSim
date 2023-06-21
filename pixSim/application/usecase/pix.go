package usecase

import (
	"errors"

	"github.com/alllga/pixSim/domain/model"
)

type PixUseCase struct {

	PixKeyRepository model.PixKeyRepositoryInterface
}

func (pixCase *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := pixCase.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, key, kind)
	if err != nil {
		return nil, err
	}
	
	pixCase.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New("unable to create new pix key at the moment")
	}

	return pixKey, nil
}

func (pixCase *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := pixCase.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}