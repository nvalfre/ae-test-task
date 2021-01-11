package memory

import (
	"ae-test-task/domain"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type AccountStoreInterface interface {
	InsertAccount(account *domain.Account) error
	InsertNewTransaction(transaction *domain.TransactionBody) error
	FindAccountByID(string) (*domain.TransactionBody, error)
	FindTransactionByID(string) (*domain.TransactionBody, error)
	FetchTransactions() (map[string]*domain.TransactionBody, error)
}

var uniqueAccount string

type AccountStoreService struct {
	db *DB
}

func NewAccountStoreService(db *DB) *AccountStoreService {
	return &AccountStoreService{db: db}
}

func (s *AccountStoreService) InsertAccount(account *domain.Account) error{
	if _, ok := s.db.accounts[account.ID]; ok {
		return errors.New("account already exist")
	}
	s.db.accounts[account.ID] = account
	uniqueAccount = account.ID
	return nil
}

func (s *AccountStoreService) InsertNewTransaction(transaction *domain.TransactionBody) error{
	if _, ok := s.db.transactions[transaction.ID]; ok {
		return errors.New("transaction already exist")
	}
	s.db.transactions[transaction.ID] = transaction
	logrus.WithFields(logrus.Fields{
		"file":    "memory_service",
		"service": "insert_transaction",
		"method":  "InsertNewTransaction",
		"message":  fmt.Sprintf("Transaction inserted successfully, transaction: %s", transaction.ID),
	})
	s.insertTransactionIntoAccount(transaction)
	return nil
}

func (s *AccountStoreService) insertTransactionIntoAccount(transaction *domain.TransactionBody) {
	account := s.db.accounts[uniqueAccount]
	account.TransactionMovements = append(account.TransactionMovements, *transaction)
	account.Balance.ImpactNewTransaction(*transaction)
	s.db.accounts[uniqueAccount] = account
	logrus.WithFields(logrus.Fields{
		"file":    "memory_service",
		"service": "insert_transaction_on_account",
		"method":  "InsertNewTransactionIntoAccount",
		"message": fmt.Sprintf("Transaction inserted successfully, transaction: %s, account: %s", transaction.ID, uniqueAccount),
	})
}

func (s *AccountStoreService) FindAccountByID(id string) (*domain.Account, error){
	if account, ok := s.db.accounts[id]; ok {
		return account, nil
	}
	return nil, errors.New("account not found")
}

func (s *AccountStoreService) FindTransactionByID(id string) (*domain.TransactionBody, error){
	if transaction, ok := s.db.transactions[id]; ok {
		return transaction, nil
	}
	return nil, errors.New("transaction not found")
}

func (s *AccountStoreService) FetchTransactions() (map[string]*domain.TransactionBody, error){
	return s.db.transactions, nil
}