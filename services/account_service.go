package services

import (
	"ae-test-task/app/memory"
	"ae-test-task/domain"
	"ae-test-task/services/builder"
	"ae-test-task/services/validator"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type AccountServiceInterface interface {
	CreateAccount(user string) (*domain.Account, error)
	GetAccount() (*domain.Account, error)
	InsertTransaction(req *domain.TransactionReq) (*domain.TransactionBody, error)
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

func (s *AccountService) GetAccount() (*domain.Account, error) {
	logrus.WithFields(logrus.Fields{
		"file":    "account_service",
		"service": "search",
		"method":  "GetAccount",
	})
	account, err := s.Store.GetAccount()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "search",
			"method":  "GetAccount",
			"err":     err,
			"message": "cannot get account",
		})
		return nil, err
	}
	return account, nil
}
func (s *AccountService) InsertTransaction(req *domain.TransactionReq) (*domain.TransactionBody, error) {
	logrus.WithFields(logrus.Fields{
		"file":    "account_service",
		"service": "insert",
		"method":  "InsertTransaction",
	})

	err := s.TransactionValidator.Validate(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "validate",
			"method":  "InsertTransaction",
			"error":   err,
		})
		return nil, err
	}

	account, err := s.GetAccount()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "get_account",
			"method":  "InsertTransaction",
			"error":   err,
		})
		return nil, err
	}

	err = s.AccountValidator.ValidateTransactionOnAccount(account, req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "validate",
			"method":  "validate_account",
			"error":   err,
		})
		return nil, err
	}

	trUUID := uuid.New()
	date := fmt.Sprintf("%s", time.Now())
	transaction := domain.TransactionBody{
		ID:            fmt.Sprintf("%s", trUUID),
		Amount:        req.Amount,
		CardType:      req.CardType,
		EffectiveDate: &date,
	}
	err = s.Store.InsertNewTransaction(&transaction)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_service",
			"service": "insert",
			"method":  "InsertTransaction",
			"err":     err,
			"message": "cannot insert new req",
		})
		return nil, err
	}
	return &transaction, err
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

func (s *AccountService) parseUUID(uuid uuid.UUID) string {
	return fmt.Sprintf("%s", uuid)
}
