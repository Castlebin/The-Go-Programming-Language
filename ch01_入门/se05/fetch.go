package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil { // 如果出错
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // 打印错误信息
			os.Exit(1)                                 // 退出程序
		}

		b, err := ioutil.ReadAll(resp.Body) // 读取响应体
		resp.Body.Close()                   // 关闭响应体
		if err != nil {                     // 如果出错
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err) // 打印错误信息
			os.Exit(1)                                                  // 退出程序
		}

		// 打印响应体
		fmt.Printf("%s", b)
	}

}
