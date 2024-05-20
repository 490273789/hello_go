package main

import "fmt"

func main() {
	createMap()
	optionMap()
	// map的相关操作

}

func createMap() {
	// map[type]type  键值对
	a := make(map[int]string, 10) // map可以存放10个键值对
	a[2009542] = "邵楠"
	a[2009573] = "少真"
	fmt.Println(a)
	// map集合在使用前一定要make
	// map的key-value是无序的
	// key是不可以重复的，后一个value会替换前一个value

	// map的创建方式
	// 方式2：
	b := make(map[int]string)
	b[2023110801] = "1号"
	b[2023110892] = "2号"
	fmt.Println(b)

	// 方式3：
	c := map[int]string{
		2023110801: "1号",
		2023110802: "2号",
	}
	c[2023110803] = "3号"
	fmt.Println(c)
}

func optionMap() {
	a := make(map[int]string)
	a[2023110801] = "w"
	a[2023110802] = "s"
	a[2023110803] = "n"
	// 删除
	delete(a, 2023110801)

	// 清空
	value, flag := a[2023110802]
	value1, flag1 := a[2023110801]

	fmt.Println(a)
	fmt.Println(value)
	fmt.Println(flag)
	fmt.Println(value1)
	fmt.Println(flag1)
	// map的长度
	fmt.Println(len(a))
	// map的遍历
	for k, v := range a {
		fmt.Println(k, v)
	}

	b := make(map[string]map[int]string)
	b["班级1"] = make(map[int]string, 3)
	b["班级1"][1330110217] = "sb"
	b["班级1"][1330110218] = "sz"
	b["班级1"][1330110219] = "aq"
	b["班级2"] = make(map[int]string, 3)
	b["班级2"][1330110220] = "cn"
	b["班级2"][1330110221] = "yx"
	b["班级2"][1330110222] = "wz"

	for k, v := range b {
		for ck, cv := range v {
			fmt.Println("班级:", k, "学号:", ck, "姓名:", cv)
		}
	}

	fmt.Println(b)
}
