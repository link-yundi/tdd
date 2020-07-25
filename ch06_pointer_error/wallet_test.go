package ch06_pointer_error

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Bitcoin int

// 定义Bitcoin的String()
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// 定义私有变量
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func TestDepositWithdraw(t *testing.T) {
	wallet := Wallet{Bitcoin(10)}
	assert := assert.New(t)
	t.Run("Deposit", func(t *testing.T) {
		wallet.Deposit(Bitcoin(10))
		//t.Logf("%s", wallet.balance)
		assert.Equal(Bitcoin(20), wallet.Balance())
	})
	t.Run("Withdraw", func(t *testing.T) {
		if err := wallet.Withdraw(Bitcoin(30)); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(Bitcoin(10), wallet.Balance())
		}
	})
}
