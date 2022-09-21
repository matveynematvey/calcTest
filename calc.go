package calc

import (
	"errors"
)

var ErrResNotInt = errors.New("result is not integer")
var ErrZero = errors.New("divided by zero")

func Add(a int, b int) (int, error) {
	return a + b, nil
}

func Div(a int, b int) (int, error) {
	if b == 0 {
		return 0, ErrZero
	}
	if (a % b) != 0 {
		return 0, ErrResNotInt
	}
	return a / b, nil
}
