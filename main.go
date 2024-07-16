package actionsplayground

import (
	"errors"
	"fmt"
)

// var ErrFoo = errors.New("error")

func returnsErr() error {
	return errors.New("error")
	// return ErrFoo
}

func main() {
	fmt.Println("Hello world")
}
