package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println("年：", now.Year())
	fmt.Println("月：", now.Month())
	fmt.Println("月：", int(now.Month()))
	fmt.Println("日：", now.Day())
	fmt.Println("时：", now.Hour())
	fmt.Println("分：", now.Minute())
	fmt.Println("秒：", now.Second())

	// Sprintf 可以将拼接好的字符串返回
	dateTime := fmt.Sprintf("时间: %d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateTime)

	formatTime := now.Format("2006-01-02 15:04:05")
	fmt.Println(formatTime)
}
