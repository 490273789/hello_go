package oop

import "fmt"

func CallBase() {
	definedStruck()
	conversionStruct()
}

type Developer struct {
	Name string
	Age  int
	Jop  string
}

func definedStruck() {
	// 创建方式1
	var d1 Developer
	d1.Name = "wsn"
	d1.Age = 30
	d1.Jop = "frond end"

	fmt.Println("d1:", d1)

	var d2 = Developer{
		Name: "wsn",
		Age:  20,
		Jop:  "back end",
	}
	d2.Age = 30

	fmt.Println("d2:", d2)

	// d3为实例的指针
	var d3 = new(Developer)
	(*d3).Name = "wsz"
	(*d3).Age = 23
	d3.Jop = "all end"

	fmt.Println("d3:", d3, *d3)

	// d4也为指针
	d4 := &Developer{"aq", 45, "go end"}

	fmt.Println("d4:", d4, *d4)

}

type Student struct {
	Age int
}

type Person struct {
	Age int
}

// 结构体之间的转换
func conversionStruct() {
	// 结构体单独定义类型，和其他类型转换时需要有完全相同的名字
	var s Student = Student{10}
	var p Person = Person{15}
	s = Student(p)
	fmt.Println(s)
	fmt.Println(p)

}

// 结构体进行type重新定义（相当于取别名），Golang认为是新的数据类型，但是相互间可以强转

type Book struct {
	name string
}

type Boo Book

// Boo是一个新的类型了

// 方法是作用在指定的数据类型上和指定的数据类型绑定，因此自定义类型都可以有方法
