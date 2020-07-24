package ch06_pointer_error

type Wallet struct {
	// 定义私有变量
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
