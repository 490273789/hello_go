package main

import "fmt"

func main() {

	// 切片是go中特有的数据类型
	// 切片是一种建立在数组类型之上的抽象，切片是对数组一个连续片段的引用，切片是一个引用类型。这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。
	// 需要注意的是，终止索引标识的项不包括在切片内。切片内提供了一个相关数组的动态窗口。
	DefineSlice()
}

func DefineSlice() {
	// 定义数组
	var intArr = [6]int{3, 6, 9, 1, 4, 7}
	// 切片建立在数组之上
	// 定义一个切片名字为slice, []动态变化的数组长度不写， int类型
	// [1:3] 将intArr数组从索引1开始，3结束（不包含3） - [1,3)
	// var slice []int = intArr[1:3]
	slice := intArr[1:3]
	fmt.Println(slice)
	fmt.Println("切片元素的个数：", len(slice))
	fmt.Println("切片的容量：", cap(slice))

	// 切片是有3个字段的数据结构（结构体），一个是指向底层数组的指针，一个是切片的长度，一个是切片的容量
	// make函数创建切片：1.切片类型 2.切片的长度 3.切片容量
	slice1 := make([]int, 4, 20)
	fmt.Println(slice1)

	slice2 := []int{1, 4, 7}
	fmt.Println(slice2)

	// 切片的遍历和数组的遍历是一样的
}
