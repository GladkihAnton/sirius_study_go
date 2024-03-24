package interfaces

import (
	"fmt"
	"reflect"
)

type namedGeometry struct {
	name string
	geometry
}

func name1(temp any) {
	//var value int
	fmt.Println(reflect.TypeOf(temp))
	intTemp, ok := temp.(int)
	if !ok {
		fmt.Println(intTemp)
	}

	switch intTemp := temp.(type) {
	case int:
		square := Square{intTemp, intTemp}
		fmt.Println(square)
	}
}

func MainEmb() {
	square := Square{1, 2}
	x := namedGeometry{"asdasd", square}
	fmt.Println(x)

	name1(1)
	fmt.Println("AFTER INT")
	name1("1")
}
