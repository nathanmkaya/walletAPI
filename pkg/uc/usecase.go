package uc

import "walletAPI/pkg/entity"

type AccountUsecase interface {
	CheckBalance(Id int) (balance float64, err error)
	MiniStatement(Id int) (Transactions []entity.Transaction, err error)
	CreateAccount(a entity.Account) (int, error)
	GetAccountById(Id int) (account *entity.Account, err error)
}

type TransactionUsecase interface {
	MakeDeposit(a *entity.Account, amount float64) error
	MakeWithdrawal(a *entity.Account, amount float64) error
}
