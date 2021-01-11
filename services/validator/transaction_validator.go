package validator

import (
	"ae-test-task/domain"
	"errors"
)

type TransactionValidatorService struct{}

type TransactionValidatorInterface interface {
	Validate(transaction *domain.TransactionBody) error
}

func (v *TransactionValidatorService) Validate(transaction *domain.TransactionReq) error {
	if err := validateAmount(transaction); err != nil {
		return err
	}
	if err := validateTransactionType(transaction); err != nil {
		return err
	}
	return nil
}

func validateAmount(transaction *domain.TransactionReq) error {
	if transaction.IsNegativeBalance() {
		return errors.New(domain.NegativeAmountError)
	}
	return nil
}

func validateTransactionType(transaction *domain.TransactionReq) error {
	if transaction.IsNotValidType() {
		return errors.New(domain.InvalidTypeError)
	}
	return nil
}
