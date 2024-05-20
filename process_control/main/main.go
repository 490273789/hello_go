package main

import "fmt"

func main() {
	// return 结束整个函数
	// if语句
	// go中 if 后面的条件可以不加(), 但是必须要加一个空格
	// if后面可以定义变量
	if count := 30; count > 0 {
		println(count)
	}

	//forLoop()
	//forRange()
	gotoUse()

}

func switchProcess() {
	//switch
	// case后面不需要加break，不会穿透

	score := 90
	switch score / 10 {
	case 1, 2, 3, 4:
		println("不及格")
	case 9:
		println("优秀")
	default:
		println("一般般")
	}

	// switch后不加任何东西，当做if用
	a := 1
	switch {
	case a == 1:
		println("a is 1")
	case a == 2:
		println("a is 2")
	}

	// 特殊用法2：switch后写变量定义，用;号结尾
	switch b := 2; {
	case b == 1:
		println("a is 1")
	case b == 2:
		println("a is 2")
	}

	// switch 穿透 - fallthrough
	c := 90
	switch c / 10 {
	case 1, 2, 3, 4:
		println("不及格")
	case 9:
		println("优秀")
		fallthrough
	case 8:
		println("还不错")
		fallthrough
	case 10:
		println("神")
	default:
		println("一般般")
	}
}

func forLoop() {
	// for中初始化变量不能用var
	// ()可以省略
	// 循环可以加标签（js也可以）
	// break 可以终结离她最近的循环
	// continue 跳出本次循环，继续下次循环

loop1:
	for i := 1; i <= 5; i++ {
		for j := 0; j <= 4; j++ {
			if j == 2 {
				continue
			}
			println(i, j)

			if i == 3 && j == 3 {
				break loop1 // 结束指定标签的循环
			}
		}
	}

}

func forRange() {
	str := "hello 邵楠"
	for i, value := range str {
		fmt.Printf("i is %d, value is %c \n", i, value)
	}
}

func gotoUse() {
	// goto 可以无条件的转移到程序的指定的行
	// 一般配合条件语句是用
	println("is 1")
	println("is 2")
	goto label
	println("is 3")
	println("is 4")
label:
	println("is 5")
	println("is 6")
}
