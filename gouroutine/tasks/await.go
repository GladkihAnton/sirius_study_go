package tasks

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func AwaitTask() {
	t := func() any {
		return atoi("12312")

	}
	result := await(t)
	fmt.Println(result, reflect.TypeOf(result))
}

func await(fn func() any) any {
	resultCh := make(chan any)

	go func(resultCh chan any) {
		resultCh <- fn()
	}(resultCh)

	return <-resultCh
}

func atoi(str string) int {
	v, _ := strconv.Atoi(str)
	time.Sleep(time.Second * 1)
	return v
}
