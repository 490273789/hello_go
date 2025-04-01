package channel

import (
	"fmt"
	"time"
)

func CallRecover() {
	// 多个协程工作，其中一个出现panic，导致程序崩溃
	// 解决办法：利用refer + recover捕获panic进行处理，即使协程出现问题，主线程仍然不受影响继续执行
	// 案例：
	go devide()
	go printNumber()
	time.Sleep(time.Second * 5)
}

func printNumber() {
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func devide() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("device报错：", err)
		}
	}()
	num1 := 10
	num2 := 0

	result := num1 / num2
	fmt.Println("devide:", result)
}
