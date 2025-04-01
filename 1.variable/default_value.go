package variable

import "fmt"

// 各种数据类型的默认值
func DefaultValue() {

	var i int
	fmt.Printf("default value of int: %d \n", i)

	var b byte
	fmt.Printf("default value of byte: %d \n", b)

	var f float32
	fmt.Printf("default value of float32: %.2f \n", f)

	var t bool
	fmt.Printf("default value of bool: %t \n", t)

	var s string
	fmt.Printf("default value of string: %#v \n", s)

	var r rune
	fmt.Printf("default value of rune: %d [%c] %#v\n", r, r, r)

	var arr []int
	fmt.Printf("default value of rune: %v %#v, arr is nil: %t \n", r, r, arr == nil)
}
