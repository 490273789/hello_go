package channel

import "fmt"

func CallBase() {
	// 1.管道的本质是一个数据结构-队列
	// 2. 数据是先进先出
	// 3. 自身线程安全，多协程访问时，不需要加锁，channel本身是线程安全的
	// 4. 管道有类型，一个string的管道只能存放string类型的数据
	// 5. 定义管道：var 变量名 chan 数据类型
	// 6. 数据类型指的是管道类型，里面放入数据的类型，管道是有类型的，intChan只能写入整数int
	// 7. 管道是引用类型，必须初始化才能写入，即mark后才能使用
	baseUse()
}

func baseUse() {

	// 定义管道
	// 使用make初始化chan， 容量为3
	intChan := make(chan int, 3)

	// 管道是引用类型的
	fmt.Printf("管道是引用类型的：%v \n", intChan)

	// 向管道存放数据
	intChan <- 10
	num := 20
	intChan <- num
	fmt.Printf("管道的长度为：%v， 管道的容量为：%v \n", len(intChan), cap(intChan))

	// 在管道中都数据
	num1 := <-intChan
	num2 := <-intChan
	// num3 := <-intChan 	//在没有协程的情况下，如果管道的数据已经全部取出，那么就会报错

	fmt.Println(num1)
	fmt.Println(num2)
	// fmt.Println(num3)

	// 管道关闭
	// 使用内置函数close可以关闭管道，当管道关闭后，就不能在向管道写数据了，但是仍然可以从该管道读取数据
	close(intChan)

	stringChan := make(chan string, 5)
	stringChan <- "w"
	stringChan <- "s"
	stringChan <- "n"
	stringChan <- "帅"
	stringChan <- "b"
	// 在遍历前如果没有关闭管道，就会出现deadlock的错误
	// 遍历管道 for - range
	close(stringChan)
	for value := range stringChan {
		fmt.Println("遍历：", value)
	}

	// 只写管道
	intChanWrite := make(chan<- int, 2)
	intChanWrite <- 10
	intChanWrite <- 20

	// 只读管道
	intChanRead := make(<-chan int, 3)
	fmt.Println(intChanRead)

}
