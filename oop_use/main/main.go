package main

import (
	"fmt"
	"hello_go/oop"
)

func main() {
	// 创建person结构体实例
	p := oop.NewPerson("王绍楠")
	p.SetAge(30)
	fmt.Println("姓名：", p.Name)
	fmt.Println("年龄：", p.GetAge())
}
