package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(float64(0))) // "8"

	fmt.Println(unsafe.Sizeof(x)) // "32" (8+16+8)

	fmt.Println(unsafe.Alignof(x))    // "8"
	fmt.Println(unsafe.Offsetof(x.a)) // "0"
	fmt.Println(unsafe.Offsetof(x.b)) // "2"
	fmt.Println(unsafe.Offsetof(x.c)) // "8"
}

var x struct {
	a bool
	b int16
	c []int
}
