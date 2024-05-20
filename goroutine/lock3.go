package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var wgLock3 sync.WaitGroup
var RWLock sync.RWMutex // 护持锁

// 多个协程操作一个数据，因为是并行的，所以数据是乱的，解决方案，互斥锁
func CallLock3() {
	wgLock3.Add(5)
	for i := 0; i < 4; i++ {
		go read()
	}
	go write()
	wgLock3.Wait()
	fmt.Println("num:", num2)

	wgLock3.Wait()
}

func read() {
	defer wgLock3.Done()
	RWLock.RLock() // 如果只是都数据这个锁不会产生影响，如果是写数据就会起作用，会等待写入锁完事在读
	fmt.Println("读取数据")
	time.Sleep(time.Second)
	fmt.Println("都数据完成")
	RWLock.RUnlock()
}

func write() {
	defer wgLock3.Done()
	RWLock.Lock()
	fmt.Println("写入数据")
	time.Sleep(time.Second * 10)
	fmt.Println("写入完成")
	RWLock.Unlock()
}
