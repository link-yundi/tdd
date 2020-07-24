package ch06

import (
	. "github.com/link-yundi/tdd/ch06_pointer_error"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	want := 10
	assert.Equal(t, want, got)
}
