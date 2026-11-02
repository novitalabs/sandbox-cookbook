package main

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2,3) = %d, want 5", got)
	}
}

func TestSub(t *testing.T) {
	if got := Sub(7, 4); got != 3 {
		t.Fatalf("Sub(7,4) = %d, want 3", got)
	}
}

func TestMul(t *testing.T) {
	if got := Mul(6, 7); got != 42 {
		t.Fatalf("Mul(6,7) = %d, want 42", got)
	}
}

func TestDiv(t *testing.T) {
	q, err := Div(10, 2)
	if err != nil || q != 5 {
		t.Fatalf("Div(10,2) = %d, %v; want 5, nil", q, err)
	}
	if _, err := Div(1, 0); !errors.Is(err, ErrDivideByZero) {
		t.Fatalf("Div(1,0) err = %v, want ErrDivideByZero", err)
	}
}
