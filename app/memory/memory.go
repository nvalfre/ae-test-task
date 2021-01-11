package memory

import "ae-test-task/domain"

type DB struct {
	accounts     map[string]*domain.Account
	transactions map[string]*domain.TransactionBody
}

func New() *DB {
	return &DB{
		accounts:     make(map[string]*domain.Account),
		transactions: make(map[string]*domain.TransactionBody),
	}
}
