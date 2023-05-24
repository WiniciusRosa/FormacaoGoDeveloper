package main

import (
	"fmt"
	"time"
)

// Erro
type MeuError struct {
	When time.Time
	What string
}

func (e MeuError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MeuError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"O arquivo sumiuuuu",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
