package entity

import "time"

type TransactionType int

const (
	Deposit TransactionType = iota + 1
	Withdraw
)

type Transaction struct {
	Id int
	TransactionType
	TransactionTime time.Time
	Amount          float64
}
