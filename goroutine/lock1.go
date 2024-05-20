package goroutine

import (
	"fmt"
	"sync"
)

var wgLock sync.WaitGroup
var num int

// 多个协程操作一个数据，因为是并行的，所以数据是乱的，解决方案，互斥锁
func CallLock1() {
	wgLock.Add(2)
	go add()
	go reduce()
	wgLock.Wait()
	fmt.Println("num:", num)
}

func add() {
	for i := 0; i < 1000; i++ {
		num += 1
	}
	wgLock.Done()
}

func reduce() {
	for i := 0; i <= 1000; i++ {
		num -= 1
	}
	wgLock.Done()
}
