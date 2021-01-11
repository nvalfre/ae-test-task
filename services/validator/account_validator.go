package validator

import (
	"ae-test-task/domain"
)

type AccountValidatorService struct{}

type AccountValidatorInterface interface {
	Validate(account *domain.Account) error
}

func (v *AccountValidatorService) Validate(account *domain.Account) error {
		if err := validateUser(account); err != nil {
			return err
		}
		if err := validateUUID(account); err != nil {
			return err
		}
		if err := validateBalance(account); err != nil {
			return err
		}
	return nil
}

func validateBalance(account *domain.Account) error {
	return nil
}

func validateUUID(account *domain.Account) error {
	return nil
}

func validateUser(account *domain.Account) error {
	return nil
}
