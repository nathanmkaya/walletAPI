package main

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"log"
	"net/http"
	"walletAPI/pkg/entity"
	"walletAPI/pkg/http/rest"
	"walletAPI/pkg/repo"
	"walletAPI/pkg/repo/postres"
	"walletAPI/pkg/uc"
)

func main() {
	var (
		ac     uc.AccountUsecase
		tx     uc.TransactionUsecase
		acRepo repo.AccountRepository
		//txRepo repo.TransactionRepository
	)

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "root",
		Database: "postgres",
	})
	defer db.Close()

	err := createSchema(db, true)
	if err != nil {
		panic(err)
	}

	acRepo = postres.NewSQLAccountRepo(db)

	ac = uc.NewAccountUsecase(acRepo)
	tx = uc.NewTransactionUsecase(acRepo)

	router := rest.Handler(ac, tx)

	fmt.Println("The server is running on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
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
