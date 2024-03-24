package my_string

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	expected := "a"
	input := []string{"a"}

	result := LongestCommonPrefix(input)
	if expected != result {
		t.Fatalf("expected: %v result: %v", expected, result)
	}
}
