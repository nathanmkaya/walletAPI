package postres

import (
	"github.com/go-pg/pg/v9"
	"walletAPI/pkg/entity"
)

type pgTransactionRepo struct {
	DB *pg.DB
}

func (p *pgTransactionRepo) Fetch(num int, account *entity.Account) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	err := p.DB.Model(account.Transactions).Limit(num).Select()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (p *pgTransactionRepo) GetByID(id int) (*entity.Transaction, error) {
	transaction := &entity.Transaction{Id: id}
	err := p.DB.Select(transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (p *pgTransactionRepo) Create(t *entity.Transaction) (int, error) {
	err := p.DB.Insert(t)
	if err != nil {
		return 0, err
	}
	return t.Id, nil
}

func (p *pgTransactionRepo) Update(t *entity.Transaction) (*entity.Transaction, error) {
	err := p.DB.Update(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (p *pgTransactionRepo) Delete(id int) (bool, error) {
	transaction := &entity.Transaction{Id: id}
	err := p.DB.Delete(transaction)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewSQLTransactionRepo(db *pg.DB) *pgTransactionRepo {
	return &pgTransactionRepo{DB: db}
}
