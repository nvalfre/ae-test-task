package services

import (
	"ae-test-task/app/memory"
	"ae-test-task/domain"
	"ae-test-task/services/builder"
	"ae-test-task/services/validator"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AccountServiceInterface interface {
	CreateAccount(user string) (*domain.Account, error)
	InsertTransaction(*domain.TransactionBody) (*domain.TransactionBody, error)
	TransactionHistory() (map[string]*domain.TransactionBody, error)
	FindTransactionByID(string2 string) (*domain.TransactionBody, error)
}
type AccountService struct {
	AccountBuilderService builder.AccountBuilderService
	AccountValidator      validator.AccountValidatorService
	TransactionValidator  validator.TransactionValidatorService
	Store                 *memory.AccountStoreService
}

func (s *AccountService) CreateAccount(user string) (*domain.Account, error) {
	accountUUID, err := uuid.NewUUID()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "create",
			"method":  "create_account",
			"error":   err,
		})
		return nil, err
	}
	userID := user
	accountID := s.parseUUID(accountUUID)
	account, err := s.AccountBuilderService.BuildNewAccount(accountID, userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "create",
			"method":  "create_account",
			"error":   err,
		})
		return nil, err
	}
	err = s.AccountValidator.Validate(account)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "validate",
			"method":  "validate_account",
			"error":   err,
		})
		return nil, err
	}
	return account, err
}

func (s *AccountService) parseUUID(uuid uuid.UUID) string {
	return fmt.Sprintf("%s", uuid)
}
func (s *AccountService) InsertTransaction(transaction *domain.TransactionBody) (*domain.TransactionBody, error) {
	logrus.WithFields(logrus.Fields{
		"file":    "account_service",
		"service": "insert",
		"method":  "InsertTransaction",
	})
	err := s.Store.InsertNewTransaction(transaction)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "insert",
			"method":  "InsertTransaction",
			"err":     err,
			"message": "cannot insert new transaction",
		})
		return nil, err
	}
	err = s.TransactionValidator.Validate(transaction)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "validate",
			"method":  "validate_transaction",
			"error":   err,
		})
		return nil, err
	}
	return transaction, err
}

func (s *AccountService) TransactionHistory() (map[string]*domain.TransactionBody, error) {
	logrus.WithFields(logrus.Fields{
		"file":    "account_service",
		"service": "history",
		"method":  "TransactionHistory",
	})
	accountTransactionHistory, err := s.Store.FetchTransactions()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "history",
			"method":  "TransactionHistory",
			"err":     err,
			"message": "cannot get transaction history",
		})
		return nil, err
	}
	return accountTransactionHistory, err
}

func (s *AccountService) FindTransactionByID(id string) (*domain.TransactionBody, error) {
	logrus.WithFields(logrus.Fields{
		"file":    "account_service",
		"service": "search",
		"method":  "FindTransactionByID",
	})
	transaction, err := s.Store.FindTransactionByID(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "search",
			"method":  "FindTransactionByID",
			"err":     err,
			"message": "cannot get transaction",
		})
		return nil, err
	}
	return transaction, nil
}
