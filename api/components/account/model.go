package account

import (
	"account_system/api/api_errors"
	"account_system/api/common"
	"account_system/api/components/transaction"
	"account_system/api/components/user"
	"sync"
)

type Account struct {
	Id           common.IdType
	User         user.User
	Balance      common.AmountType
	Transactions []transaction.Transaction
	Lock         sync.RWMutex
}

func New() *Account {
	a := &Account{}
	a.Id = common.NewId()
	a.User = *user.New("Juan Perez")
	return a
}

func (a *Account) Debit(amount common.AmountType) (*transaction.Transaction, *api_errors.ApiError) {
	a.Lock.Lock()
	if a.Balance < amount {
		a.Lock.Unlock()
		return nil, api_errors.InsufficientFundsError()
	}
	a.Balance -= amount
	newT := transaction.New(amount, transaction.DEBIT)
	a.Transactions = append(a.Transactions, *newT)
	a.Lock.Unlock()
	return transaction.Copy(newT), nil
}

func (a *Account) Credit(amount common.AmountType) (*transaction.Transaction, *api_errors.ApiError){
	a.Lock.Lock()
	a.Balance += amount
	newT := transaction.New(amount, transaction.CREDIT)
	a.Transactions = append(a.Transactions, *newT)
	a.Lock.Unlock()
	return transaction.Copy(newT), nil
}

func (a *Account) GetBalance() common.AmountType {
	a.Lock.RLock()
	balance := a.Balance
	a.Lock.RUnlock()
	return balance
}

func (a *Account) GetTransactions() []transaction.Transaction {
	a.Lock.RLock()
	transactions := make([]transaction.Transaction, len(a.Transactions))
	copy(transactions, a.Transactions)
	a.Lock.RUnlock()
	return transactions
}

func (a *Account) GetTransaction(tId common.IdType) (*transaction.Transaction, *api_errors.ApiError) {
	a.Lock.RLock()
	var newT *transaction.Transaction
	for _, t := range a.Transactions {
		if tId == t.Id {
			newT = transaction.Copy(&t);
		}
	}
	a.Lock.RUnlock()
	if newT == nil {
		return nil, api_errors.TransactionNotFoundError()
	}
	return newT, nil
}

func createMainAccount() *Account {
	return New()
}
