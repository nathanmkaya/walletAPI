package uc

import (
	"errors"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
	"github/nathanmkaya/walletAPI/mocks"
	"github/nathanmkaya/walletAPI/pkg/entity"
	"testing"
)

func TestNewTransactionUsecase(t *testing.T) {
	assert.Implements(t, (*TransactionUsecase)(nil), &ucTransactionUsecase{})
}

func TestUcTransactionUsecase_MakeDeposit(t *testing.T) {
	var account *entity.Account
	f := fuzz.New().NilChance(0)
	f.Fuzz(&account)
	mockAccountRepo := new(mocks.AccountRepository)

	amount := 20.0
	account.Balance = account.Balance + amount

	mockAccountRepo.On("Update", account).Return(nil, nil)

	usecase := NewTransactionUsecase(mockAccountRepo)

	err := usecase.MakeDeposit(account, amount)
	assert.Nil(t, err)

	mockAccountRepo.AssertExpectations(t)
}

func TestUcTransactionUsecase_MakeDepositError(t *testing.T) {
	var account *entity.Account
	f := fuzz.New().NilChance(0)
	f.Fuzz(&account)
	mockAccountRepo := new(mocks.AccountRepository)

	amount := 20.0
	account.Balance = account.Balance + amount

	mockAccountRepo.On("Update", account).Return(nil, errors.New("error occurred"))

	usecase := NewTransactionUsecase(mockAccountRepo)

	err := usecase.MakeDeposit(account, amount)
	assert.NotNil(t, err)

	mockAccountRepo.AssertExpectations(t)
}

func TestUcTransactionUsecase_MakeWithdrawal(t *testing.T) {
	var account *entity.Account
	f := fuzz.New().NilChance(0)
	f.Fuzz(&account)
	mockAccountRepo := new(mocks.AccountRepository)

	amount := 20.0
	account.Balance = account.Balance - amount

	mockAccountRepo.On("Update", account).Return(nil, nil)

	usecase := NewTransactionUsecase(mockAccountRepo)

	err := usecase.MakeWithdrawal(account, amount)
	assert.Nil(t, err)

	mockAccountRepo.AssertExpectations(t)
}

func TestUcTransactionUsecase_MakeWithdrawalError(t *testing.T) {
	var account *entity.Account
	f := fuzz.New().NilChance(0)
	f.Fuzz(&account)
	mockAccountRepo := new(mocks.AccountRepository)

	amount := 20.0
	account.Balance = account.Balance - amount

	mockAccountRepo.On("Update", account).Return(nil, errors.New("error occurred"))

	usecase := NewTransactionUsecase(mockAccountRepo)

	err := usecase.MakeWithdrawal(account, amount)
	assert.NotNil(t, err)

	mockAccountRepo.AssertExpectations(t)
}
