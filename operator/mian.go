package main

import (
	"fmt"
)

// 变量声明
// 首字母大写，全局可以访问

var Global = "全局变量"

// 首字母小写，只能在这个包里面使用
var global = "包变量"

var (
	First        = "wsn"
	second int32 = 16
)

// 只声明不赋值，会给这个类型的默认值
var c int

// 主函数必须在 main 包中运行
func main() {
	//ff()
	//ch()
	// stringContent
	// constant()
	// lit()
	//binaryContent()

	//replaceHolder()
}
func ch() {
	var s = "hello"
	// %T type
	for _, ele := range s {
		fmt.Printf("%c %d %T\n", ele, ele, ele)
	}
}

type user struct {
	Name string
	Age  int
}

func replaceHolder() {
	u := &user{
		Name: "Tom",
		Age:  17,
	}

	fmt.Printf("v => %v \n", u)
	fmt.Printf("+v => %+v \n", u)
	fmt.Printf("#v => %#v \n", u)
	fmt.Printf("T => %T \n", u)

	// 指针: 就是一个地址，指针变量就是存储地址的变量

}

/*
二进制相关知识
*/
func binaryContent() {
	// 除法
	// 如果两个int类型相除值也是int类型的，不会出现小数
	// 取模 - a%b = a- a/b*b

	// go里面的++  -- 只能单独使用，不能参与到运算中
	// 只能写在变量的后面，不能写在变量的前面

	// 除 2 和 右移 1位是一样的
	// bb == cc
	// 乘 2 就是左移 1位
	aa := 9
	bb := aa / 2
	cc := aa >> 1
	fmt.Println(bb, cc)
	// 二进制 1 2^0  10 2^1  100 2^2 ,有n个0，2^n-1,
	// 补吗的另一种形式: 最后一个1右面部分保持不变，左边部分按位取反
	// >>
	// <<
	num1 := 256
	num2 := 1000
	num3 := 1024
	fmt.Printf("num1 - 256: %b\n", num1)  // 100000000  256 = 2^8
	fmt.Printf("num2 - 1000: %b\n", num2) // 1111101000
	fmt.Printf("num3 - 1024: %b\n", num3) // 10000000000  10个0 1024 = 2^10
	fmt.Printf("num1 - 260: %b\n", 260)   // 100000100  256 + 4 = 2^8 + 2^2 = 100000000 + 100 = 100000100
	k := 1 << 10
	m := 1 << 20
	g := 1 << 30
	fmt.Printf("%d %d %d", k, m, g)

	// & ｜ ^ 字啊react中的应用
	// & - 可以判断有无权限
	// | 添加权限
	// ^ 和 & 配合删除权限

	// 11111111
	// 无符号
	// 2^7 = 2^6 + 2^5 + 2^4 + 2^3 + 2^2 + 2^1 + 2^0 + 1
	// 128 =  64 + 32  + 16  + 8   + 4   + 2   + 1   + 1
	// 有符号
	// 2 ^ 7 - 1 = 2^6 + 2^5 + 2^4 + 2^3 + 2^2 + 2^1 + 2^0
	// 128   - 1 =  64 + 32  + 16  + 8   + 4   + 2   + 1   + 1
	// 有符号最大负数
	// 10000000
	// 减1 01111111
	// 取反 10000000  - 无符号正数 - 128
}

func lit() {
	// 字面量 - 没有出现变量名，直接出现了值，基础类型的字面量相当于是常量
	fmt.Printf("%t\n", 4 == 4.) // 4. 就是 4.0
	fmt.Printf("%v\n", .4i)
	fmt.Printf("%t\n", '\u4f17' == '众')
	fmt.Printf("Hello\n!\n")
}

// go doc 查看包的注释
// godoc 导出网页版注释文档
