package main

import (
	"fmt"
	"hello_go/type_conversion"
	"hello_go/util"
	"unsafe"
)

// 各种数据类型的默认值
func defaultValue() {
	var i int
	var b byte
	var f float32
	var t bool
	var s string
	var r rune
	var arr []int
	fmt.Printf("default value of int: %d \n", i)
	fmt.Printf("default value of byte: %d \n", b)
	fmt.Printf("default value of float32: %.2f \n", f)
	fmt.Printf("default value of bool: %t \n", t)
	fmt.Printf("default value of string: [%s] \n", s)
	fmt.Printf("default value of rune: %d [%c] \n", r, r)
	fmt.Printf("default value of rune: %v, arr is nil: %t \n", r, arr == nil)
}

func scope() {
	var i int8 = 127 // 符号位占一位 2^7 - 1, 0会占一个数 负数： -2^7
	var b byte
	fmt.Printf("default value of int: %d \n", i)
	fmt.Printf("default value of byte: %d \n", b)
}

func pointer() {
	var a int
	var pointer unsafe.Pointer = unsafe.Pointer(&a)

	var p uintptr = uintptr(pointer)
	var ptr *int = &a

	fmt.Printf("pointer %p, p %d %x, ptr %p \n", pointer, p, p, ptr)
}

type User struct {
	name string
	age  int
}

func (user User) Hello() {
	fmt.Printf("my name is %s \n", user.name)
}

//	类型别名  =
// type byte = uint8
// type rune = float32

// 自定义类型 没有 =  可以自定义成员方法
type ms map[string]int

func (user ms) Say() {
	fmt.Printf("%d \n", user["hello"])
}

type mss = map[string]int

// 下面写法是不可以的
//
//	func (user mss) Say() {
//		fmt.Printf("%d \n", user["hello"])
//	}
//
// 定义一个函数
type add func(a, b int) int

func structCont() {
	type User struct {
		name string
		age  int
	}
	user := User{"wsn", 30}
	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)

}

func useBinaryFnc() {
	var a int32 = 260
	b := util.BinaryFormat(a)
	fmt.Printf("b is %s \n", b)
}

func main() {
	// defaultValue()
	// scope()
	// pointer()
	// structCont()
	// useBinaryFnc()
	// const Self = map
	// ms.Say()
	//string_part.StringBasic()
	type_conversion.TypeConversion()
}
