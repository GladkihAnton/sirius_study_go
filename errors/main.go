package errors

import (
	"errors"
	"fmt"
)

type MyError struct {
	value string
}

func (myError MyError) Error() string {
	return fmt.Sprintf("%v", myError.value)
}

func temp() (int, error) {
	return 1, MyError{value: "test"}
}

func MainError() {
	connect, err := temp()
	if err != nil {
		//err1 := errors.As(err, &currErr)

		switch {
		case errors.Is(err, MyError{}):
		}

		fmt.Println(err.Error())
	}

	fmt.Println(connect)
}
