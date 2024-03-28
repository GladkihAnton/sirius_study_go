package errors

import (
	"errors"
	"fmt"
	"reflect"
)

var MyError2 = errors.New("asdas")
var myerror2 = errors.New("Error2")
var Temp = MyError{value: "asdasd"}
var Temp1 = MyError1{value: "Error2"}

type MyError1 struct {
	value string
	err   error
}

func (myError MyError1) Error() string {
	return fmt.Sprintf("%v", myError.value)
}

func (myError MyError1) Unwrap() error {
	return myError.err
}

type MyError struct {
	value string
	err   error
}

func (myError MyError) Error() string {
	return fmt.Sprintf("%v", myError.value)
}

func (myError MyError) Unwrap() error {
	return myError.err
}

var temp = errors.New("cannot use user")

// myerror1 - myerror - default error

func zeroTemp() (int, error) {
	value, err := firstTemp()
	return value, MyError1{"Error2", err}
}
func firstTemp() (int, error) {
	value, err := secondTemp()
	return value, MyError{"Error", err}
}

func secondTemp() (int, error) {

	return 1, temp
}

func thirdTemp() (int, error) {
	panic("Me panic")
	return 1, nil
}

func protect(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR:", err)
		} else {
			fmt.Println("Everything went smoothly!")
		}
	}()
	fn()
}

func MainError() {
	connect, err := thirdTemp()

	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("ERROR:", err1)
		} else {
			fmt.Println("Everything went smoothly!")
		}
	}()

	if err != nil {
		fmt.Println(reflect.TypeOf(err))
		fmt.Printf("%v", err)
		fmt.Println()
	}
	fmt.Printf("Success %v\n", connect)
}
