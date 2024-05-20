package main

import (
	"fmt"
	"go-base/function"
)

var age int = getAge()

func getAge() int {
	fmt.Println("全局变量定义！")
	return 10
}

func init() {
	fmt.Println("init函数执行！")
}

func main() {
	fmt.Println("main函数执行！")
	function.Call()
}
