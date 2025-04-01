package oop

import "fmt"

// 继承
// 当多个结构体存在相同的属性和方法的时候，可以从这些结构体中抽象出结构体，在该结构体中定义这些相同的属性和方法，其他的结构体不需要重新定义这些属性和方法，
// 只需要嵌套一个匿名结构体即可，也就是说如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承的特性
// 继承的优点
// 提高代码的复用性、扩展性

type Animal struct {
	Age    int
	Weight float32
}

func (animal *Animal) Shout() {
	fmt.Println("我在大叫")
}

func (animal *Animal) ShowInfo() {
	fmt.Printf("动物年龄是：%v，动物的体重是：%v", animal.Age, animal.Weight)
}

type Cat struct {
	// go中的继承，匿名结构体
	Animal
}

func (cat *Cat) scratch() {
	fmt.Println("我在挠人")
}

func CallInherit() {
	cat := &Cat{}
	cat.Animal.Age = 13
	cat.Weight = 10
	cat.Animal.Shout()
	cat.Animal.ShowInfo()
	cat.scratch()
}
