package file

import (
	"bufio"
	"fmt"
	"os"
)

func CallWriteOutFile() {
	input()
}

func input() {
	// 打开文件
	file, err := os.OpenFile("/Users/wangshaonan/code/go/go-base/file/test2.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	permiss := os.FileMode(0777).String()
	fmt.Println("权限是：", permiss)

	if err != nil {
		fmt.Printf("打开文件失败：%v \n", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, writeError := writer.WriteString("我是写入的内容1111")
	if writeError != nil {
		fmt.Printf("文件写入错误：%v \n", writeError)
		return
	}

	// 流在缓冲区，刷新数据 -- >刷新后才能写入数据
	writer.Flush()
}
