package main

import "go-base/oop"

func main() {
	// go支持面向对象的特性，但是和传统的面向对象有区别，并不是纯粹面向对象语言
	// 结构体 来实现面向对象，go中没有类
	// 我  姓名 年龄 性别
	//oop.CallMethod()
	//oop.CallBase()
	// 跨包访问结构体 - 使用工厂模式
	oop.CallInherit()

	oop.CallInterface()
}
