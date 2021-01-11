package validator

import (
	"ae-test-task/domain"
)

type TransactionValidatorService struct{}

type TransactionValidatorInterface interface {
	Validate(transaction *domain.TransactionBody) error
}

func (v *TransactionValidatorService) Validate(transaction *domain.TransactionBody) error {
	if err := validateAmount(transaction); err !=nil{
		return err
	}
	return nil
}

func validateAmount(transaction *domain.TransactionBody) error {
	return nil
}
