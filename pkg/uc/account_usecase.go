package uc

import (
	"github/nathanmkaya/walletAPI/pkg/entity"
	"github/nathanmkaya/walletAPI/pkg/repo"
)

type ucAccountUsecase struct {
	accountRepository repo.AccountRepository
}

func (accountUsecase *ucAccountUsecase) CreateAccount(a entity.Account) (Id int, err error) {
	Id, err = accountUsecase.accountRepository.Create(a)
	if err != nil {
		return 0, err
	}
	return Id, nil
}

func (accountUsecase *ucAccountUsecase) GetAccountById(id int) (account *entity.Account, err error) {
	ac, err := accountUsecase.accountRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (u *ucAccountUsecase) CheckBalance(Id int) (balance float64, err error) {
	account, err := u.accountRepository.GetByID(Id)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

func (u *ucAccountUsecase) MiniStatement(Id int) (entity.Statement, error) {
	account, err := u.accountRepository.GetByID(Id)
	if err != nil {
		return entity.Statement{}, err
	}

	return entity.Statement{
		Balance:      account.Balance,
		Transactions: account.Transactions,
	}, err
}

func NewAccountUsecase(accountRepository repo.AccountRepository) *ucAccountUsecase {
	return &ucAccountUsecase{accountRepository: accountRepository}
}
