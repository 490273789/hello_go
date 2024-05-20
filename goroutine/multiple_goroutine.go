package goroutine

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func CallMultipleGoroutine() {
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("协程-", n)
		}(i)
	}
	wg.Wait()
}
