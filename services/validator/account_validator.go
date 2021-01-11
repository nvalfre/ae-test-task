package validator

import (
	"ae-test-task/domain"
	"errors"
)

type AccountValidatorService struct{}

type AccountValidatorInterface interface {
	Validate(account *domain.Account) error
	ValidateTransactionOnAccount(account *domain.Account, req *domain.TransactionReq) error
}

func (v *AccountValidatorService) ValidateTransactionOnAccount(account *domain.Account, req *domain.TransactionReq) error {
	if err := validateRequest(account.Balance.Amount, *req.Amount, req.CardType); err != nil {
		return err
	}
	return nil
}

func validateRequest(account float64, req float64, cardType *string) error {
	switch *cardType {
	case domain.Debit:
		if (account - req) < 0 {
			return errors.New(domain.NegativeAmountError)
		}
	default:
		return nil
	}
	return nil
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
