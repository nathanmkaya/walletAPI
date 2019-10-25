package uc

import (
	"walletAPI/pkg/entity"
	"walletAPI/pkg/repo"
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

func (u *ucAccountUsecase) MiniStatement(Id int) ([]entity.Transaction, error) {
	account, err := u.accountRepository.GetByID(Id)
	if err != nil {
		return nil, err
	}
	return account.Transactions[:20], err
}

func NewAccountUsecase(accountRepository repo.AccountRepository) *ucAccountUsecase {
	return &ucAccountUsecase{accountRepository: accountRepository}
}
