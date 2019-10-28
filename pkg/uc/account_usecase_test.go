package uc

import (
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
	"github/nathanmkaya/walletAPI/mocks"
	"github/nathanmkaya/walletAPI/pkg/entity"
	"testing"
)

func TestNewAccountUsecase(t *testing.T) {
	assert.Implements(t, (*AccountUsecase)(nil), &ucAccountUsecase{})
}

func TestUcAccountUsecase_CreateAccount(t *testing.T) {
	var account entity.Account
	f := fuzz.New().NilChance(0)
	f.Fuzz(&account)
	mockAccountRepo := new(mocks.AccountRepository)

	mockAccountRepo.On("Create", account).Return(account.Id, nil)

	usecase := NewAccountUsecase(mockAccountRepo)

	id, err := usecase.CreateAccount(account)
	assert.NotNil(t, id)
	assert.Equal(t, account.Id, id)
	assert.Nil(t, err)

	mockAccountRepo.AssertExpectations(t)
}

func TestUcAccountUsecase_CheckBalance(t *testing.T) {
	account := &entity.Account{
		Id:      99,
		Balance: 100,
	}
	mockAccountRepo := new(mocks.AccountRepository)

	mockAccountRepo.On("GetByID", account.Id).Return(account, nil)

	usecase := NewAccountUsecase(mockAccountRepo)

	balance, err := usecase.CheckBalance(account.Id)
	assert.Equal(t, account.Balance, balance)
	assert.Nil(t, err)

	mockAccountRepo.AssertExpectations(t)
}

func TestUcAccountUsecase_GetAccountById(t *testing.T) {
	var account *entity.Account
	f := fuzz.New().NilChance(0)
	f.Fuzz(&account)
	mockAccountRepo := new(mocks.AccountRepository)

	mockAccountRepo.On("GetByID", account.Id).Return(account, nil)

	usecase := NewAccountUsecase(mockAccountRepo)

	ac, err := usecase.GetAccountById(account.Id)
	assert.Nil(t, err)
	assert.Equal(t, account, ac)

	mockAccountRepo.AssertExpectations(t)
}

func TestUcAccountUsecase_MiniStatement(t *testing.T) {
	var account *entity.Account
	f := fuzz.New().NilChance(0)
	f.Fuzz(&account)
	mockAccountRepo := new(mocks.AccountRepository)

	mockAccountRepo.On("GetByID", account.Id).Return(account, nil)

	usecase := NewAccountUsecase(mockAccountRepo)

	statement, err := usecase.MiniStatement(account.Id)
	assert.Nil(t, err)
	assert.Equal(t, account.Balance, statement.Balance)
	assert.ElementsMatch(t, account.Transactions, statement.Transactions)

	mockAccountRepo.AssertExpectations(t)
}
