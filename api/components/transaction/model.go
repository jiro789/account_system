package transaction

import (
	"account_system/api/common"
)

type Type string

const (
	DEBIT  Type = "debit"
	CREDIT Type = "credit"
)

type Transaction struct {
	Id            common.IdType     `json:"id"`
	Type          Type              `json:"type"`
	Amount        common.AmountType `json:"amount"`
	EffectiveDate common.DateType   `json:"effective_date"`
}

func New(amount common.AmountType, transactionType Type) *Transaction {
	t := &Transaction{}
	t.Id = common.NewId()
	t.Type = transactionType
	t.Amount = amount
	t.EffectiveDate = common.NewDate()

	return t
}

func Copy(other *Transaction) *Transaction {
	newT := &Transaction{
		Id:            other.Id,
		Amount:        other.Amount,
		Type:          other.Type,
		EffectiveDate: other.EffectiveDate,
	}
	return newT
}
