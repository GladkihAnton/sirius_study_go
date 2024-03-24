package pointers

import "unsafe"

func Main() {
	//a := 8 // asdasdasd
	//var b *int
	// & взять ссылку от значение
	//  *взять значение от ссылки
	mainTask1()
	//changeIntegerByPointer(&a)

	//b := [3]int{1, 2, 3}
	//changeArray(b)

	//fmt.Println(a)
	//fmt.Println(b)
}

func changeIntegerByPointer(a *int) {
	*a = 9
}

func changeArray(a [3]int, pointer unsafe.Pointer) {
	a[0] = 10
}
