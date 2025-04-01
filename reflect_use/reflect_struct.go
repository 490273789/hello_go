package reflect_use

import (
	"fmt"
	"reflect"
)

func CallRefluctStruct() {
	stu := Student{Name: "帅王", Age: 18}
	useReflect1(stu)
}

type Student struct {
	Name string
	Age  int
}

func useReflect1(i interface{}) {
	// 1. 调用TypeOf函数，返回reflect.Type类型的数据
	reflectType := reflect.TypeOf(i)
	fmt.Println("reflectType:", reflectType)
	fmt.Printf("reflectType的具体类型是： %T: \n", reflectType)
	// 2. 调用ValueOf函数，返回reflect.Value类型数据
	reflectValue := reflect.ValueOf(i)
	fmt.Println("reflectValue:", reflectValue)
	fmt.Printf("reflectValue的具体类型是： %T: \n", reflectValue)

	result := reflectValue.Interface()
	origin, flage := result.(Student)
	if flage {
		fmt.Println(origin)
	}
}
