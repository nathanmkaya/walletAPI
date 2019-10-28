package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github/nathanmkaya/walletAPI/pkg/entity"
	"github/nathanmkaya/walletAPI/pkg/helpers"
	"github/nathanmkaya/walletAPI/pkg/uc"
	"log"
	"net/http"
	"strconv"
)

func Handler(ac uc.AccountUsecase, tx uc.TransactionUsecase) http.Handler {
	router := httprouter.New()

	router.POST("/account", CreateAccount(ac))
	router.GET("/account/:id", GetAccount(ac))
	router.GET("/account/:id/balance", CheckBalance(ac))
	router.POST("/account/:id/withdraw", Withdraw(tx, ac))
	router.POST("/account/:id/deposit", Deposit(tx, ac))
	router.GET("/account/:id/statement", GetMiniStatement(ac))

	return router
}

func GetAccount(usecase uc.AccountUsecase) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Account ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}
		account, err := usecase.GetAccountById(Id)
		if err != nil {
			http.Error(w, fmt.Sprintf("The Account: #%s does not exist.", p.ByName("id")), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(account)
	}
}

func CreateAccount(usecase uc.AccountUsecase) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var account entity.Account
		err := helpers.DecodeJSONBody(w, r, &account)
		if err != nil {
			var mr *helpers.MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		_, err = usecase.CreateAccount(account)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode("New Account created.")
	}
}

func CheckBalance(usecase uc.AccountUsecase) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Account ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}
		balance, err1 := usecase.CheckBalance(Id)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(fmt.Sprintf("{'balance':%f}", balance))
	}
}

func GetMiniStatement(usecase uc.AccountUsecase) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Account ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}
		statement, err1 := usecase.MiniStatement(Id)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(statement)
	}
}

type TX struct {
	Amount float64
}

func Withdraw(usecase uc.TransactionUsecase, accountUsecase uc.AccountUsecase) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Account ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}
		var tx TX
		err = helpers.DecodeJSONBody(w, r, &tx)
		if err != nil {
			var mr *helpers.MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		account, err2 := accountUsecase.GetAccountById(Id)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return
		}
		err3 := usecase.MakeWithdrawal(account, tx.Amount)
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode("Withdraw Successful")
	}
}

func Deposit(usecase uc.TransactionUsecase, accountUsecase uc.AccountUsecase) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Account ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}
		var tx TX
		err = helpers.DecodeJSONBody(w, r, &tx)
		if err != nil {
			var mr *helpers.MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		account, err2 := accountUsecase.GetAccountById(Id)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return
		}
		err3 := usecase.MakeDeposit(account, tx.Amount)
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode("Deposit Successful")
	}
}
