package repo

import (
	"walletAPI/pkg/entity"
)

//go:generate mockery -all -output $PWD/mocks

type AccountRepository interface {
	Fetch(num int) ([]*entity.Account, error)
	GetByID(id int) (*entity.Account, error)
	Create(p entity.Account) (int, error)
	Update(p *entity.Account) (*entity.Account, error)
	Delete(id int) (bool, error)
}

type TransactionRepository interface {
	Fetch(num int) ([]*entity.Transaction, error)
	GetByID(id int) (*entity.Transaction, error)
	Create(p entity.Transaction) (int, error)
	Update(p *entity.Transaction) (*entity.Transaction, error)
	Delete(id int) (bool, error)
}
