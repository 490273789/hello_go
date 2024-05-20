package main

import "fmt"

func main() {
	name := "w"
	// 1. 在内存中为 name 变量开辟了一个空间，这个空间存储的值为 "w"
	// 2. 这个空间本身是有地址的， & + 变量 就可以获取变量的地址(可以想象为旅店，每个房间都有房间号)
	fmt.Printf("name的内存地址：%v \n", &name) // 地址 0x1400008e260

	// 指针变量 - 指针就是内存地址，就是一个存储指针的变量
	// 1.定义一个指针变量
	var ptr *string = &name
	// 2. var代表要声明一个指针变量
	// 3. ptr指针变量的名字
	// 4. *string 变量的类型（可理解为指向string类型的指针）
	// 5. &age是一个地址，是ptr变量具体的值
	fmt.Printf("ptr变量的值：%v \n", ptr)
	fmt.Printf("ptr变量的指针地址：%v \n", &ptr)
	// 6. 获取指针地址对应的值，*+变量名 访问指针指向的空间
	fmt.Printf("ptr变量存储的值(也就是存的指针)指向的值： %s \n", *ptr)
	// 7. 通过指针改变指向的值
	var age = 18
	fmt.Printf("age的地址：%v \n", &age)
	// var pointerAge = &age
	// *pointerAge = 20
	// 上面两行简写为下面1
	*&age = 20
	fmt.Printf("*&age的值：%v \n", age)
	fmt.Printf("age的地址：%v \n", &age)
	age = 30
	fmt.Printf("age的值：%v \n", age)
	fmt.Printf("age的地址：%v \n", &age)

	// 指针变量接收的一定是个地址
	// 指针变量的地址不可以不匹配，举例：不能把*int变量的地址赋值给*float类型的指针
	// 标识(zhi)符
	// 变量名，方法名，只要是起名字的地方的那个名字都叫标识符  var age int = 9  age就是标识符

	// 键盘输入
	var name1 string
	fmt.Println("请输入年龄")
	_, err := fmt.Scanln(&name1)
	if err != nil {
		fmt.Println(err)
	}

}
