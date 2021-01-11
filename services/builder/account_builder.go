package builder

import (
	"ae-test-task/app/memory"
	"ae-test-task/domain"
	"ae-test-task/services/validator"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type AccountBuilderService struct {
	Store     *memory.AccountStoreService
	Validator validator.AccountValidatorService
}

type AccountBuilderServiceInterface interface {
	BuildNewAccount(name string, userID string) (*domain.Account, error)
	create(account *domain.Account) error
}

func (s *AccountBuilderService) BuildNewAccount(uuid, userID string) (*domain.Account, error) {
	account := &domain.Account{
		ID:                   fmt.Sprintf("%s", uuid),
		User:                 &domain.UserBody{ID: &userID},
		TransactionMovements: make([]domain.TransactionBody, 5),
		Balance:              &domain.Balance{},
	}
	logrus.WithFields(logrus.Fields{
		"file":    "account_builder",
		"service": "create",
		"method":  "Create new account",
		"account": account,
	})
	err := s.create(account)
	if err != nil {
		return nil, err
	}
	return account, err
}

func (s *AccountBuilderService) create(account *domain.Account) error {
	if account.ID == "" {
		return errors.New("no account id")
	}

	s.Validator.Validate(account)

	err := s.Store.InsertAccount(account)
	return err
}
