package postres

import (
	"github.com/go-pg/pg/v9"
	"walletAPI/pkg/entity"
	"walletAPI/pkg/repo"
)

func NewSQLAccountRepo(db *pg.DB) repo.AccountRepository {
	return &pgAccountRepo{
		DB: db,
	}
}

type pgAccountRepo struct {
	DB *pg.DB
}

func (p *pgAccountRepo) Fetch(num int) ([]*entity.Account, error) {
	var accounts []*entity.Account
	err := p.DB.Model(&accounts).Limit(num).Select()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (p *pgAccountRepo) GetByID(id int) (*entity.Account, error) {
	account := &entity.Account{Id: id}
	err := p.DB.Select(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (p *pgAccountRepo) Create(a entity.Account) (int, error) {
	println("reached")
	err := p.DB.Insert(&a)
	if err != nil {
		return 0, err
	}
	return a.Id, nil
}

func (p *pgAccountRepo) Update(a *entity.Account) (*entity.Account, error) {
	err := p.DB.Update(&a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (p *pgAccountRepo) Delete(id int) (bool, error) {
	account := &entity.Account{Id: id}
	err := p.DB.Delete(account)
	if err != nil {
		return false, err
	}
	return true, nil
}
