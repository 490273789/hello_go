package function

import "fmt"

// 定义匿名函数

var Result int = func(num1 int, num2 int) int {

	return num1 + num2
}(10, 20)

// Reduce 将匿名函数赋值给变量，这个变量就是函数类型的变量
var Reduce = func(num1 int, num2 int) int {
	return num1 - num2
}

// defer

func Add(num1 int, num2 int) int {
	// 程序遇到关键字defer，不会立即执行defer后的语句，而是将defer后的语句压入一个单独的栈中，然后继续执行函数后面的语句
	// 栈的特点是先进后出
	// 在函数执行完后，从栈中取出语句开始执行
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)
	// 后面的语句不会影响上面的结果，上面类似于快照，会讲他的值也拷贝到栈中
	//应用：比如你想在开启一个功能后，要关闭，，可以在打开资源的时候，随手用defer关闭掉
	num1 += 20
	num2 += 20
	var sum int = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}
