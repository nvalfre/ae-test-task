package domain

type Balance struct {
	amount float64
}

func (b *Balance) ImpactNewTransaction(transaction TransactionBody) {
	b.amount += *transaction.Amount
}
