package tests

import (
	"fmt"
	"testing"
)

func TestMatchContains(t *testing.T) {
	var tests = []struct {
		pattern, src string
		want         bool
	}{
		{"asd", "asdaa", true},
		{"xasd", "asdaa", false},
	}

	for _, test := range tests {
		name := fmt.Sprintf("case(pattern:%v,src:%v)", test.pattern, test.src)
		t.Run(name, func(t *testing.T) {
			got := MatchContains(test.pattern, test.src)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
