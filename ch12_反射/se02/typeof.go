package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)  // t 是一个类型为 reflect.Type 的变量
	fmt.Println(t.String()) // "int"  t 实现了 String() 方法
	fmt.Println(t)          // "int"  t 实现了 String() 方法，所以也可以直接打印，和上面结果一样

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"    说明 reflect.TypeOf() 返回的是变量的实际类型，w 的类型是 *os.File

	fmt.Printf("%T\n", 3) // "int"               fmt.Printf() 的格式化参数 %T 打印变量的实际类型
	fmt.Printf("%T\n", w) // "*os.File"
}
