package goroutine

import (
	"fmt"
	"time"
)

func CallBase() {
	// 程序（program）：是为了完成特定任务，用某种语言写的一组指令的集合，是一段静态的代码。程序是静态的
	// 进程（process）：是程序的一次执行过程，正在运行的一个程序，进程作为资源分配的一个单位，在内存中回味每个进程分配不同的内存区域。
	// 进程是动态的，是一个动的过程，进程的生命周期：他有自己的产生、存在、消亡
	// 线程（thread）：进程可以进一步的细化为线程，是一个程序的一条执行路径
	// 若一个进程同一时间并行执行多个线程，就是支持多线程的。
	// 协程（goroutine）：又称为微线程，纤程，协程是一种用户态的轻量级线程
	// 作用：在执行A函数的时候，可以随时中断，去执行B函数，然后中断继续执行A函数（可以自动切换），注意这一切换过程并不是函数调用（没有调用语句）
	// 过程很像多线程，然而协程中只有一个线程在执行（协程的本质是单个线程）
	go useGoroutine()
	mainRoutine()

	// 主死从随
}

// 要求
// 1. 在主线程中开启一个goroutine，该goroutine没个1秒输出一次“帅b楠”
// 2. 在主线城中也每隔1秒输出“帅b王”，输出10次后退出程序
// 3. 要求主线程和goroutine同时执行
func useGoroutine() {
	for i := 1; i < 11; i++ {
		fmt.Println("帅B王 + ", i)
		time.Sleep(time.Second)
	}

}

func mainRoutine() {
	for i := 1; i < 11; i++ {
		fmt.Println("帅B楠 + ", i)
		time.Sleep(time.Second)
	}
}
