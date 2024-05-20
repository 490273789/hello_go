package reflect_use

import (
	"fmt"
	"reflect"
)

func CallCategory() {
	// 获取变量的类别，两种方式
	// reflect.Type.Kind()
	// reflect.Value.Kind()
	//
	person := Person{Name: "帅王", Age: 18}

	useReflect2(&person)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Print() {
	fmt.Println("人的姓名：", p.Name)
	fmt.Println("人的年龄：", p.Age)
}

func (p *Person) Set(name string, age int) {
	fmt.Println("Call Set", name, age)
	p.Name = name
	p.Age = age
}

func (p Person) GetSum(n1, n2 int) int {
	fmt.Println("Call Set", n1, n2)
	return n1 + n2
}

func useReflect2(i interface{}) {
	// 调用ValueOf函数，返回reflect.Value类型数据
	reflectValue := reflect.ValueOf(i)

	// NumMethod 获取方结构体内方法的数量（只计算方法体名字是大写的）
	n2 := reflectValue.Elem().NumMethod()
	println("结构体方法的数量：", n2)
	var params []reflect.Value
	params = append(params, reflect.ValueOf("帅男"))
	params = append(params, reflect.ValueOf(10))
	reflectValue.Method(2).Call(params)

	// NumFiled 获取结构体内部有多少项
	n1 := reflectValue.Elem().NumField()
	reflectValue.Elem().Field(0).SetString("帅少")

	println("结构体中变量的数量：", n1)

	for i := 0; i < n1; i++ {
		// fmt.Printf("第%d个字段是: %v \n", i, reflectValue.Field(i)) // 值传递
		fmt.Printf("第%d个字段是: %v \n", i, reflectValue.Elem().Field(i)) // 引用值传递

	}

}
