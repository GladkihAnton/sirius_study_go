package my_string

import (
	"fmt"
	"strings"
)

func Main() {
	a := "rз"
	b := "rз"

	y := fmt.Sprintf("%v %v", a, b)
	fmt.Println(y)

	fmt.Println("")
	builder := strings.Builder{}
	builder.WriteString(a)
	builder.WriteString(" ")
	builder.WriteString(a)
	x := builder.String()

	fmt.Println(x)

	for _, value := range a {
		fmt.Printf("%c %v", value, value)
	}

}
