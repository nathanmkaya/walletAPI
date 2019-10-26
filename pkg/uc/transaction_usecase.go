package uc

import (
	"time"
	"walletAPI/pkg/entity"
	"walletAPI/pkg/repo"
)

type ucTransactionUsecase struct {
	transactionRepo repo.AccountRepository
}

func (u *ucTransactionUsecase) MakeDeposit(a *entity.Account, amount float64) (err error) {
	transaction := entity.Transaction{
		TransactionType: entity.Deposit,
		Amount:          amount,
		TransactionTime: time.Now(),
	}
	transactions := a.Transactions
	a.Transactions = append(transactions, transaction)
	a.Balance = a.Balance + amount
	_, err = u.transactionRepo.Update(a)
	if err != nil {
		return err
	}
	return nil
}

func (u *ucTransactionUsecase) MakeWithdrawal(a *entity.Account, amount float64) (err error) {
	transaction := entity.Transaction{
		TransactionType: entity.Withdraw,
		Amount:          amount,
		TransactionTime: time.Now(),
	}
	transactions := a.Transactions
	a.Transactions = append(transactions, transaction)
	a.Balance = a.Balance - amount
	_, err = u.transactionRepo.Update(a)
	if err != nil {
		return
	}
	return nil
}

func NewTransactionUsecase(repository repo.AccountRepository) TransactionUsecase {
	return &ucTransactionUsecase{transactionRepo: repository}
}
