package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CallBase() {
	readFile()
	fileStream1()
	fileStream2()
}

func readFile() {
	// 打开文件
	file, err := os.Open("/Users/wangshaonan/code/go/hello_go/file/test.txt")
	if err != nil {
		fmt.Println("文件打开错误：", err)
	}
	// 没有错误
	fmt.Printf("文件=%v \n", file)

	// 关闭文件
	closeErr := file.Close()
	if closeErr != nil {
		println("关闭失败：", closeErr)
	} else {
		println("文件关闭成功！")
	}
}

// 读取文件的内容并显示在终端（使用io util一次将整个文件读入到内存中），这种方式适用于文件不大的情况，相关方法和函数os.ReadFile
func fileStream1() {
	// 输入流/输出流：针对程序来定义方向
	// 流向程序的 - 输入流 - 读取
	// 流向目标文件的 - 输出流 - 写出

	content, err := os.ReadFile("/Users/wangshaonan/code/go/hello_go/file/test.txt")
	if err != nil {
		println("文件读取错误：", err)
	}

	fmt.Printf("文件内容-ASIIC码为：%v \n", content)
	fmt.Printf("文件内容为：%v \n", string(content))
}

// fileStream2 读取文件的内容并显示在终端（带缓冲区的方式），适合读取比较大的文件，使用os.Open,file.Close,bufio.NewReader(),reader.ReadString函数和方法
func fileStream2() {
	file, err := os.Open("/Users/wangshaonan/code/go/hello_go/file/test1.txt")
	if err != nil {
		fmt.Println("文件打开错误：", err)
	}
	//defer file.Close()
	// 创建一个流
	stream := bufio.NewReader(file)

	for {
		str, err := stream.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}

	closeErr := file.Close()
	if closeErr != nil {
		fmt.Println("文件读取失败：", closeErr)
	}

	fmt.Println("文件读取成功，并且全部读取完成！")
}
