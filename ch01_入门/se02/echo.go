// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// os.Args 是一个字符串（string）的切片（slice）, 存放命令行参数，其中 os.Args[0] 是命令本身的名字.随后的是命令行参数
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " " // 拼接命令行参数
	}
	// 打印命令行参数
	fmt.Println(s)
}
