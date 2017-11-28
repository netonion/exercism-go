package account

import (
	"sync"
)

type Account struct {
	sync.Mutex
	balance int64
	closed  bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}

func (acc *Account) Close() (int64, bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed {
		return 0, false
	}

	acc.closed = true
	return acc.balance, true
}

func (acc *Account) Balance() (int64, bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed {
		return 0, false
	}

	return acc.balance, true
}

func (acc *Account) Deposit(amount int64) (int64, bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed {
		return 0, false
	}

	if amount < 0 && acc.balance+amount < 0 {
		return 0, false
	}

	acc.balance = acc.balance + amount
	return acc.balance, true
}
