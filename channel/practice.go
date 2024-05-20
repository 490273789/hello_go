package channel

import (
	"fmt"
	"sync"
	"time"
)

// 1. 开启一个writeData协程，向管道中写50个数
// 2. 开启一个readData协程，从管道中读取writeData写入的数据
// 3. 注意：writeData和readData操作的是同一个管道
// 4. 主线程需要等待writeData和readData协程都完成后才能退出
var wg sync.WaitGroup

func CallPractice() {
	intChan := make(chan int, 10)
	wg.Add(2)
	go writeChan(intChan)
	go readChan(intChan)
	wg.Wait()
}

func writeChan(intChan chan int) {
	defer wg.Done()
	fmt.Println("开始写入数据")
	for i := 0; i < 10; i++ {
		intChan <- i
		fmt.Println("写入数据：", i)
		time.Sleep(time.Second)
	}
	close(intChan)
	fmt.Println("完成写入数据，关闭管道")
}

func readChan(intChan chan int) {
	defer wg.Done()
	for value := range intChan {
		fmt.Println("读取数据：", value)
	}
}

// 当管道只写不读，就会出现管道阻塞的问题
//
