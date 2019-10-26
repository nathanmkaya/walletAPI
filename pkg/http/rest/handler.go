package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"walletAPI/pkg/entity"
	"walletAPI/pkg/uc"
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
		json.NewEncoder(w).Encode(account)
	}
}

func CreateAccount(usecase uc.AccountUsecase) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decode := json.NewDecoder(r.Body)
		var account entity.Account
		err := decode.Decode(&account)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = usecase.CreateAccount(account)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New Account created.")
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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fmt.Sprintf("{'balance':%f}", balance))
	}
}

func GetMiniStatement(usecase uc.AccountUsecase) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Account ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}
		transactions, err1 := usecase.MiniStatement(Id)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transactions)
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
		decode := json.NewDecoder(r.Body)
		var tx TX
		err1 := decode.Decode(&tx)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		account, err2 := accountUsecase.GetAccountById(Id)
		if err2 != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err3 := usecase.MakeWithdrawal(account, tx.Amount)
		if err3 != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Withdraw Successful")
	}
}

func Deposit(usecase uc.TransactionUsecase, accountUsecase uc.AccountUsecase) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Account ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}
		decode := json.NewDecoder(r.Body)
		var tx TX
		err1 := decode.Decode(&tx)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		account, err2 := accountUsecase.GetAccountById(Id)
		if err2 != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err3 := usecase.MakeDeposit(account, tx.Amount)
		if err3 != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Deposit Successful")
	}
}
