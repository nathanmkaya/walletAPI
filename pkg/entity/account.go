package entity

type Account struct {
	Id           int
	Balance      float64
	Transactions []Transaction
}
