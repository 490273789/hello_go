package reflect_use

import (
	"fmt"
	"reflect"
)

func CallReflect() {
	// 对基本类型数据进行反射
	// 定义一个基本数据类型
	var num int = 100
	useReflect(num)
}

// 空接口可以接收任何类型
func useReflect(i interface{}) {
	// 1. 调用TypeOf函数，返回reflect.Type类型的数据
	reflectType := reflect.TypeOf(i)
	fmt.Println("reflectType:", reflectType)
	fmt.Printf("reflectType的具体类型是： %T: \n", reflectType)
	// 2. 调用ValueOf函数，返回reflect.Value类型数据
	reflectValue := reflect.ValueOf(i)
	fmt.Println("reflectValue:", reflectValue)
	fmt.Printf("reflectValue的具体类型是： %T: \n", reflectValue)

	// 调用Int方法，返回v持有的有符号整数
	num2 := 80 + reflectValue.Int()
	fmt.Println(num2)
}
