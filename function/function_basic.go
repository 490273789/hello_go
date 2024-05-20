package function

import "fmt"

// 函数定义：为完成某功能的程序指令的集合

/**
func 函数名(形参列表) 返回值类型列表 {
	执行语句
	return +返回值列表
}

形参：接受外来的数据
实参：实际掺入的数据
返回值数量：0个 到 n个
如果返回值只有一个，返回值类型的括号可以不写
*/

func calc(num1 int, num2 int) int {
	var sum int = 0
	return sum + num1 + num2
}

func calc1(num1 int, num2 int) (int, int) {
	var sum int = 0
	return sum + num1 + num2, num1 - num2
}

// 可变参数, 不传args是一个空数组
func calc2(args ...int) int {
	fmt.Println("可选参数：", args)
	var sum int = 0
	for _, value := range args {
		sum += value
	}
	return sum
}

// 基本数据类型和数组默认都是值传递，即进行值拷贝，在函数内修改，不会影响到原来的值

// 以值传递方式的数据类型，如果希望在函数内部修改函数外部的变量，可以传入变量的地址，函数内以指针的方式操作变量，从效果上来看类似引用传递

// 在go中函数也是一种数据类型，可以赋值给一个变量，则该变量就是函数类型的变量了

// 函数可以作为参数传递给其他函数

func calc3(num1 int, num2 *int) {
	num1 = 3
	*num2 = 3
}

// 自定义数据类型 - 给已存在的类型起别名
// type 自定义数据类型名 自定义数据类型
type myInt int

// 函数返回值进行命名
func calc4(num1 int, num2 int) (res1 int, res2 int) {
	res1 = num1 + num2
	res2 = num1 - num2
	return
}

func Call() {
	//sum := calc(10, 20)
	//fmt.Println(sum)
	//sum1, sum2 := calc1(10, 20)
	//fmt.Println(sum1)
	//fmt.Println(sum2)
	//// 如果只想接收1个参数，不想接收的参数是用 _ 匿名变量接收
	//sum3, _ := calc1(30, 20)
	//fmt.Println(sum3)
	//
	//sum4 := calc2()
	//sum5 := calc2(1, 2, 3, 4)
	//fmt.Println("calc2:", sum4, sum5)

	sum6 := 2
	sum7 := 2
	calc3(sum6, &sum7)
	fmt.Printf("calc3 called: sum6 is %d, sum7 is %d \n", sum6, sum7)

	calc4(20, 30)
	fmt.Print("calc4 called: ")
}
