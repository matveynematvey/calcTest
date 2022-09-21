package calc

import (
	"testing"
)

var (
	res int
	err error
)

func TestAdd(t *testing.T) {
	res, err = Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if res != 3 {
		t.Error("sum expected to be 3")
	}
}

func TestDiv(t *testing.T) {
	res, err = Div(1, 0)
	if err != ErrZero {
		t.Error("unexpected error")
	}

	res, err = Div(3, 2)
	if err != ErrResNotInt {
		t.Error("unexpected error")
	}

	res, err = Div(4, 2)
	if res != 2 {
		t.Error("unexpected error")
	}

}
