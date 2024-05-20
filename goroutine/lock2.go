package goroutine

import (
	"fmt"
	"sync"
)

var wgLock2 sync.WaitGroup
var lock sync.Mutex // 护持锁
var num2 int

// 多个协程操作一个数据，因为是并行的，所以数据是乱的，解决方案，互斥锁
func CallLock2() {
	wgLock2.Add(2)
	go add2()
	go reduce2()
	wgLock2.Wait()
	fmt.Println("num:", num2)
}

func add2() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		num += 1
		lock.Unlock()
	}
	wgLock2.Done()
}

func reduce2() {
	for i := 0; i <= 1000; i++ {
		lock.Lock()
		num -= 1
		lock.Unlock()
	}
	wgLock2.Done()
}
