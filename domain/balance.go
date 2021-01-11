package domain

type Balance struct {
	Amount float64
}

func (b *Balance) ImpactNewTransaction(transaction TransactionBody) {
	if *transaction.CardType == Debit {
		b.Amount -= *transaction.Amount
	}
	if *transaction.CardType == Credit {
		b.Amount += *transaction.Amount
	}
}
