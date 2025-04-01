package oop

import "fmt"

func CallMethod() {
	useMethod()
}

type A struct {
	Num int
}

type B struct {
	Name string
}

// 方法 注意和函数的区别

func (a A) getA() int {
	return a.Num
}

// 结构体对象传入方法中是值传递
func (a A) set() {
	a.Num = 20
}

// 传递指针
func (b *B) setPointer() {
	b.Name = "wsn"
	fmt.Println("this is setPointer:", b)
}

type cStr string

func (str cStr) print() {
	fmt.Println(str)
}

// 如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量String()进行输出
type Company struct {
	Name    string
	Address string
}

/*
如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量String()进行输出
建议定义结构的时候直接定义好String方法
*/
func (company *Company) String() string {
	str := fmt.Sprintf("Name = %v, Age = %v", company.Name, company.Address)
	return str
}
func useMethod() {
	// 规则：方法的访问范围控制和函数一样，方法名首字母小写，只能在本包访问，方法首字母大写，可以在本包和其他包访问

	var a = A{10}
	num := a.getA()
	fmt.Println(num)
	a.set()
	fmt.Println("this is a:", a)
	// 传递指针
	// 注意不推荐一个结构值传递和指针传递同时使用
	b := B{"wsz"}
	(&b).setPointer()
	fmt.Println("this is b:", b)

	var str cStr
	str = "qwe"
	str.print()

	company := Company{"六度云", "成府路"}
	//fmt.Println(company)
	fmt.Println(&company) // 注意如果是引用只需要加&

	str = "wsn"

}
