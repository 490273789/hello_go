package file

import (
	"fmt"
	"os"
)

func CallCopy() {
	copy()
}

func copy() {
	// 源文件
	originFilePath := "/Users/wangshaonan/code/go/hello_go/file/test2.txt"
	// 目标文件
	targetFilePath := "/Users/wangshaonan/code/go/hello_go/file/test3.txt"

	// 读取文件
	content, readErr := os.ReadFile(originFilePath)

	if readErr != nil {
		fmt.Printf("文件读取失败：%v \n", readErr)
	}

	writeErr := os.WriteFile(targetFilePath, content, 0666)
	if writeErr != nil {
		fmt.Printf("写入错误：%v \n", writeErr)
	}
}
