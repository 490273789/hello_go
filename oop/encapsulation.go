package oop

import "fmt"

// 封装
// 封装（encapsulation）就是抽象出的字段和对字段的操作封装在一起，数据被保护在内部，程序的其他包只有通过被授权的操作方法，才能对字段进行操作
// 简单说就是让数据的访问和修改更加程序化，更加合理，将数据进行保护惊醒限制，让访问更加安全合理
// 好处：
// 1、隐藏细节
// 2、可以对数据进行验证，保证安全合理

// 实现封装
// 建议将结构体、字段的首字母小写（其他包不能使用，类似private，实际开发不小写也可以，因为封装没那么严格）
// 给结构体所在的包提供一个工厂模式的函数，首字母大写
//提供一个首字母大写的set方法，（类似其他语言的public），用于对属性判断并赋值
// 提供一个首字母大写的get方法用于获取属性的值

type person struct {
	Name string
	age  int // 其他包不能直接进行访问
}

// NewPerson 定义工厂函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// SetAge 定义set和get方法，对age字段进行封装，因为在函数中可以加一些列的限制操作，确保被封装字段的安全合理性
func (person *person) SetAge(age int) {
	if age > 0 && age < 150 {
		person.age = age
	} else {
		fmt.Println("对不起，您传入的年龄范围不正确！")
	}
}

func (person *person) GetAge() int {
	return person.age
}
