package actionsplayground

import (
	"errors"
	"fmt"
)

var ErrFoo = errors.New("error")

func returnsErr() error {
	return ErrFoo
}

func main() {
	fmt.Println("Hello world")
}
