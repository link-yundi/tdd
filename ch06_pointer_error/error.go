package ch06_pointer_error

import "errors"

var ErrInsufficientFunds error = errors.New("cannot withdraw, insufficient funds.")
