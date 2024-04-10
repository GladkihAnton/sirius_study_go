package gouroutine

import (
	"fmt"
	"go.uber.org/goleak"
	"testing"
)

func TestGoroutineLeak(t *testing.T) {
	defer goleak.VerifyNone(t)

	MainGoroutine()
	fmt.Println("YPA")
}
