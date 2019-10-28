package postres

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github/nathanmkaya/walletAPI/pkg/entity"
	"github/nathanmkaya/walletAPI/pkg/repo"
	"testing"
)

type AccountRepoTestSuite struct {
	suite.Suite
	db         *pg.DB
	account    entity.Account
	repository repo.AccountRepository
}

func (suite *AccountRepoTestSuite) SetupSuite() {
	suite.db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "root",
		Database: "postgres",
	})
	_ = createSchema(suite.db, true)

	f := fuzz.New().NilChance(0)
	f.Fuzz(&suite.account)
	suite.repository = NewSQLAccountRepo(suite.db)
}

func (s *AccountRepoTestSuite) TearDownSuite() {
	_ = s.db.Close()
}

func createSchema(db *pg.DB, testing bool) error {
	for _, model := range []interface{}{(*entity.Account)(nil), (*entity.Transaction)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: testing,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func TestAccountRepositoryImplementation(t *testing.T) {
	assert.Implements(t, (*repo.AccountRepository)(nil), &pgAccountRepo{})
}

func (s *AccountRepoTestSuite) TestPgAccountRepo_Create() {
	t := s.T()
	assert.NotNil(t, s.db)

	// Test creating a new account
	assert.NotNil(t, s.account)
	id, err := s.repository.Create(s.account)
	assert.NotEqual(t, id, 0)
	assert.Nil(t, err)

	// Test creating an existing account
	id2, err2 := s.repository.Create(s.account)
	assert.Equal(t, id2, 0)
	assert.NotNil(t, err2)
}

func (s *AccountRepoTestSuite) TestPgAccountRepo_Fetch() {
	t := s.T()
	createdAccountId, _ := s.repository.Create(s.account)
	account, err := s.repository.Fetch(1)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	createdAccount, _ := s.repository.GetByID(createdAccountId)

	assert.Contains(t, account, createdAccount)
}

func (s *AccountRepoTestSuite) TestPgAccountRepo_GetByID() {
	t := s.T()
	account1 := entity.Account{
		Id:      99,
		Balance: 100,
	}
	//var in = &s.account
	_, _ = s.repository.Create(account1)
	account, err := s.repository.GetByID(account1.Id)
	assert.Nil(t, err)
	assert.NotNil(t, account)

	assert.Equal(t, account1, *account)
}

func (s *AccountRepoTestSuite) TestPgAccountRepo_Update() {
	t := s.T()
	_, _ = s.repository.Create(s.account)
	account, err := s.repository.GetByID(s.account.Id)
	assert.Nil(t, err)
	assert.NotNil(t, account)

	account.Balance = 20
	account_updated, err1 := s.repository.Update(account)
	assert.Nil(t, err1)
	assert.NotNil(t, account_updated)

	assert.NotEqual(t, s.account, account_updated)

}

func (s *AccountRepoTestSuite) TestPgAccountRepo_Delete() {
	t := s.T()
	_, _ = s.repository.Create(s.account)
	result, err := s.repository.Delete(s.account.Id)
	assert.Nil(t, err)
	assert.Equal(t, result, true)

	result1, err1 := s.repository.Delete(s.account.Id)
	assert.NotNil(t, err1)
	assert.Equal(t, result1, false)

	account, err1 := s.repository.GetByID(s.account.Id)
	assert.Nil(t, account)
	assert.NotNil(t, err1)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(AccountRepoTestSuite))
}
