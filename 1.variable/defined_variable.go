package variable

import "fmt"

// 变量的定义
func DefinedVariable() {
	// 方式1： 先声明， 在赋值
	var name0 string
	fmt.Printf("name0: %s \n", name0)
	name0 = "王"
	fmt.Printf("name0: %s \n", name0)

	// 方式2：直接赋值
	var name1 = "邵"
	fmt.Printf("name1: %s \n", name1)

	// 方式3：使用短声明符号 :=  这种方式只能在函数中使用，不能在全局使用
	name2 := "楠"
	fmt.Printf("name2: %s \n", name2)

	// 方式4：定义多个变量
	var name3, name4 = "小王", "小赵"
	fmt.Printf("name3: %s \nname4: %s \n", name3, name4)

	// 方式5：定义常量 - 必须要直接赋值，复制后不能修改
	// 有类型常量 & 无类型常量
	const name5 = "楠啊" // 无类型常量，使用更加灵活简洁
	const num int = 1  // 有类型常量
	fmt.Printf("name5: %s \n", name5)
	fmt.Printf("num: %d \n", num)
}
