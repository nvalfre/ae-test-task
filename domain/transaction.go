package domain

type TransactionBody struct {
	ID            string   `json:"id"`
	CardType      *string  `json:"card_type"`
	Amount        *float64 `json:"Amount"`
	EffectiveDate *string  `json:"effective_date"`
}

type TransactionReq struct {
	CardType *string  `json:"card_type"`
	Amount   *float64 `json:"Amount"`
}

func (r *TransactionReq) IsNotValidType() bool {
	return *r.CardType != Debit && *r.CardType != Credit
}

func (r *TransactionReq) IsNegativeBalance() bool {
	return *r.Amount < 0
}
