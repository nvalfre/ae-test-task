package domain

type Balance struct {
	amount float64
}

func (b *Balance) ImpactNewTransaction(transaction TransactionBody) {
	if *transaction.CardType == Debit {
		b.amount -= *transaction.Amount
	}
	if *transaction.CardType == Credit {
		b.amount += *transaction.Amount
	}
}
