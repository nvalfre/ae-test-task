package domain

type TransactionBody struct {
	ID            string  `json:"id"`
	CardType      *string  `json:"card_type"`
	Amount        *float64 `json:"amount"`
	EffectiveDate *string  `json:"effective_date"`
}
