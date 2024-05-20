package oop

import "fmt"

func CallInterface() {
	c := &Chinese{}

	a := American{}

	greet(c)
	greet(a)

	test()
}

// SayHello 接口的定义：定义规则、定义规范、定义某种能力
// SayHello 接口可以只声明，不用实现
type SayHello interface {
	sayHello()
}

// 接口的实现
// 中国人

type Chinese struct {
}

// 实现接口的方法 --- 具体的实现

func (person Chinese) sayHello() {
	fmt.Println("你好")
}

func (person Chinese) sayChinese() {
	fmt.Println("说中文")
}

type American struct {
}

// 实现接口的方法 --- 具体的实现

func (person American) sayHello() {
	fmt.Println("Hello")
}

// 定义一个函数，专门用来打招呼的函数，接收具备SayHelloo接口能力的变量
func greet(s SayHello) {
	s.sayHello()
	// 断言
	//ch, flag := s.(Chinese) // 看s能否转成Chinese类型并且赋给ch变量
	//if flag == true {
	//	ch.sayChinese()
	//}
	ch, flag := s.(*Chinese)
	fmt.Println("s.(Chinese):", ch, flag)
	if ch, flag := s.(*Chinese); flag {
		ch.sayChinese()
	}

	switch s.(type) { // type属于go中的一个关键字，固定写法
	case Chinese:
		fmt.Println("Chinese")
	case American:
		fmt.Println("American")
	}
}

// 多态：变量（实例）具有多种状态。面向对象的第三大特征，在go中，多态特征是通过接口实现的。可以按照统一的接口来调用不同的实现（同一个接口不同的实现）。这时接口的变量就呈现不同的形态
// 多态数组
func test() {
	//var arr [2]SayHello
	//arr[0] = Chinese{}
	//arr[1] = American{}
	//fmt.Println(arr)
}

// 断言
//Go中语法中有一个可以直接判断是否是该类型的变量，value, ok = element.(T),这里value就是变量的只，ok是bool类型，element是interface变量 T是断言类型
