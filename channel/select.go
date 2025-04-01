package channel

import (
	"fmt"
	"time"
)

func CallSelect() {
	// select: 解决多个管道的选择问题，也可以叫做多路复用，可以从多个管道中随机公平的选择一个来执行
	// ps：case后面必须进行的是io操作，不能是等值，随机选择一个IO操作
	// ps：default防止select被阻塞住，加入default

	intChan := make(chan int, 2)
	go func() {
		time.Sleep(time.Second * 3)
		intChan <- 10
		intChan <- 20
	}()
	stringChan := make(chan string, 2)
	go func() {
		time.Sleep(time.Second * 2)
		stringChan <- "w"
		stringChan <- "s"
	}()

	// fmt.Println(<-intChan) // 本身取数据就是阻塞的
	select {
	case v := <-intChan:
		fmt.Println("intChan:", intChan, v)
	case v := <-stringChan:
		fmt.Println("stringChan:", stringChan, v)
	default:
		fmt.Println("防止阻塞")
	}

}
