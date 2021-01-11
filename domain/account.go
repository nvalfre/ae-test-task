package domain

type Account struct {
	ID                   string            `json:"id"`
	User                 *UserBody          `json:"user_id"`
	TransactionMovements []TransactionBody `json:"transaction_movements"`
	Balance              *Balance           `json:"balance"`
}
