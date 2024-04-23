package main

import (
	"errors"
	"fmt"
)

var ErrRecordNotFound = errors.New("record not found")
var ErrDuplicateFound = errors.New("duplicate")

func fuga() error {
	return fmt.Errorf("fuga: %w", ErrRecordNotFound)
}

type ErrValidate struct {
	Err error
}

func (e *ErrValidate) Error() string {
	return e.Err.Error()
}

func main() {
	err := fuga()
	if err != nil {
		if validateErr, ok := err.(*ErrValidate); ok {
			fmt.Println("ErrRecordNotFound:", validateErr.Err)
		}

		if errors.Is(err, ErrRecordNotFound) {
			fmt.Println("ErrRecordNotFound:", err)
		} else if errors.Is(err, ErrDuplicateFound) {
			fmt.Println("ErrDuplicateFound:", err)
		}
	}
}
