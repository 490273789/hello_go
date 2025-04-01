package main // 声明文件所在的包，每个go文件必须有归属的包

import (
	"fmt"
	variable "hello_go/1.variable"
	"hello_go/1.variable/version"
)

// 变量种类
// 1、全局变量
// 2、局部变量

// 基本数据类型
//基本数据类型
// 数值型 - 整型， 浮点型
// 字符型 - 没有单独的字符型，是用byte来保存耽搁字母字符
// 布尔型
// 字符串

// 派生数据类型/复杂数据类型
// 指针、数组、结构体、管道、函数、切片、接口、map

// 进制
// n进制  逢n进1
// n进制转换为10进制
// 十进制转n进制
// 除n取余数，倒叙排列
// 二进制 和 十进制
// 十进制 0 1 2 3 4 5 6 7 8 9
// 八进制 0o 0 1 2 3 4 5 6 7
// 二进制 0b 0 1
// 十六进制 0x 0 1 2 33 4 5 6 7 8 9 a b c d e f

// 全局变量，其他的包可以调用，自己的包可以直接使用
var title = "学习变量"

// 变量块：定义多个变量
var (
	message     = "万丈高楼平地起"
	description = "基础最重要"
)

// API - Application Programming Interface
// 源文件以 .go 扩展名结尾
func main() {
	fmt.Printf("全局变量 title: %s \n", title)
	fmt.Printf("全局变量 message: %s \n", message)
	fmt.Printf("全局变量 description: %s \n", description)
	// 使用version包中的变量
	fmt.Printf("version: %s", version.Version)
	variable.DefinedVariable()
	variable.DefaultValue()
	scope()
}

func scope() {
	var i int8 = 127 // 符号位占一位； 2^7 - 1，0会占一个数； 负数： -2^7
	var b byte
	fmt.Printf("default value of int: %d \n", i)
	fmt.Printf("default value of byte: %d \n", b)
}

// fmt.Printf
// %s：用于字符串的占位符。
// %d：用于有符号十进制整数的占位符。
// %f：用于浮点数的占位符。
// %t：用于布尔值的占位符。
// %v：根据变量类型自动选择合适的占位符。
// %T：用于打印变量的类型。
// %c：Unicode 字符类型（rune），用于格式化输出单个字符。
// %b：二进制类型，用于格式化输出整数的二进制表示。
// %x 或 %X：十六进制类型，用于格式化输出整数的十六进制表示。
// %o：八进制类型，用于格式化输出整数的八进制表示。
// %p：指针类型，用于格式化输出指针的十六进制表示。
