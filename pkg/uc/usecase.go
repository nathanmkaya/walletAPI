package uc

import "github/nathanmkaya/walletAPI/pkg/entity"

//go:generate mockery -all -output $PWD/mocks

type AccountUsecase interface {
	CheckBalance(Id int) (balance float64, err error)
	MiniStatement(Id int) (statement entity.Statement, err error)
	CreateAccount(a entity.Account) (int, error)
	GetAccountById(Id int) (account *entity.Account, err error)
}

type TransactionUsecase interface {
	MakeDeposit(a *entity.Account, amount float64) error
	MakeWithdrawal(a *entity.Account, amount float64) error
}
