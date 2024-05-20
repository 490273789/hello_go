// 包声明： 字母和下划线组合
// 可以和文件夹名字不同
// 同一个文件夹下的声明一致
package main

import (
	"fmt"
)

// 引入包
// import ( "fmt", _ "strings")
// 如果包引入了但是没有是用，会报错
// 匿名引入：前面多了一个下划线

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

func fmtContent() {
	// 布尔
	// default: false
	// %t
	// bool: true, false

	// 整型
	// %d
	// default: 0
	// 有符号：int8,int16,int32,int64,int
	// 无符号：uint8,uint16,uint32,uint64,uint
	// 浮点型
	// default: 0
	// %f %e %g  %o %x %p
	// float32,float64

	// 指针
	// default: nil
	// %d %p
	// uintptr

	// 引用
	// default: nil
	// %v
	// map slice channel

	// 字节 - %d
	// byte，，本质是uint8

	// 任意字符 - %d
	// rune 直观理解就是字符，本质是int32，一个rune四个字节

	// 字符串 %s
	// 错误 - %v

	// fmt 格式化输出
	// 常用的：%s - 字符串占位符,%d - 数字占位符,%v,%+v,%#v
	// 常用的：%f - float占位符,%d
}

func constant() {
	// 声明常量
	const sex = "男"
	println(sex)
	// 批量声明
	const (
		PI  = 3.14
		PI1 = 3.15
	)

	const (
		PI2 = 100
		PT3 // 100跟上一行的值相同
		PI4 // 100
	)

	// iota
	const (
		v0 = iota // 默认是0
		v1        // 1
		v2        // 2
		v3        // 3
	)
	const (
		f0 = iota // 默认是0
		f1        // 1
		_         // 2
		f3        // 3
	)

	const (
		g0 = iota // 默认是0
		g1 = 30   // 1
		g2 = iota // 2
		g3        // 3
	)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
	)

	const (
		h1, h2 = iota + 1, iota + 2 // iota = 0
		h3, h4                      // iota =1
	)
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
