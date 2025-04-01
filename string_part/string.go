package string_part

import (
	"fmt"
	"strconv"
	"strings"
)

func StringBasic() {
	// go中没有专门的字符类型，如果要存储单个字符（字母），一般使用byte来保存
	// go中字符使用UTF-8编码
	stringAscII()
	// stringMethod()
	//stringJoin()
	// stringStruct()

}

// 字符类型
func stringAscII() {

}

func stringMethod() {
	//	字符串里面可以包含任意的Unicode字条
	//	\ 转义字符
	//	`` 模板字符串，里面的换行空格都不会被忽略
	//	字符串常用操作
	//  len(str) 求长度，计算的是字节的数量，比如1个汉字是3个字节那长度就是3
	// 字符串遍历
	// 1. for range循环
	// 2. 利用切片
	str0 := "王邵楠"
	r := []rune(str0)
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
	}

	// 字符串转int
	strA := "666"
	num1, _ := strconv.Atoi(strA)
	fmt.Println(num1)
	strB := strconv.Itoa(num1)
	fmt.Println(strB)

	// strings.Count() 查看字符串中有多少个子串
	// strings.EqualFold() 不去份大小写比较两个字符串，去分大小写用 ==
	// strings.replace() 字符串的替换
	// strings.ToLower
	// strings.ToUpper()
	// strings.TrimSpace() // 将左右两边的空格去除
	// strings.Trim() 将字符串左右两边指定的字符串去掉
	// strings.TrimLeft() 将字符串左边指定的字符串去掉
	// strings.TrimRight() 将字符串右边指定的字符串去掉
	//	strings.Split 分割字符串
	//	strings.Contains 判断是否包含
	//	strings.HasPrefix 前缀判断
	//	strings.HasSuffix 后缀判断
	//	strings.Index() // 子串最开始出现的位置
	//	strings.lastIndex() // 子串最后出现的位置
	str1 := "hello string type"
	arr1 := strings.Split(str1, " ") // 将字符串通过“,”分割为数组
	fmt.Printf("arr1 is %#v \n", arr1)
	fmt.Printf("str1 is contains \"str\" ? - %t \n", strings.Contains(str1, "str"))
	fmt.Printf("str1 is start with \"he\"? - %t \n", strings.HasPrefix(str1, "he"))
	fmt.Printf("str1 is end with \"pe\"? - %t \n", strings.HasSuffix(str1, "pe"))
	fmt.Printf("The first occurrence of \"ing\" in the str1 is - %d \n", strings.Index(str1, "ing"))
	fmt.Printf("The last occurrence of \"e\" in the str1 is - %d \n", strings.LastIndex(str1, "e"))
}

func stringJoin() {
	//	字符串拼接
	//	加号拼接 +  缺点每次连接新的字符串都会重新申请新的空间
	//	func fmt.Sprintf(format string, a ...interface{}) string
	//	func strings.Join(elems []string, sep string) string
	//	当有大量的string需要拼接的时候，用strings.Builder 效率更高
	str1 := "王"
	str2 := "邵"
	str3 := "楠"

	join1 := str1 + str2 + str3
	join2 := fmt.Sprintf("%s %s %s", str1, str2, str3)
	join3 := strings.Join([]string{str1, str2, str3}, " ")
	sb := strings.Builder{}
	sb.WriteString(str1)
	sb.WriteString(str2)
	sb.WriteString(str3)
	join4 := sb.String()

	fmt.Println(join1)
	fmt.Println(join2)
	fmt.Println(join3)
	fmt.Println(join4)
}

func stringStruct() {
	// 字符类型
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '('
	fmt.Println(c2)
	var c3 byte = '6'
	fmt.Println(c3)
	// 字符类型，本质上就是一个整数，也可以直接参与运算，输出字符的时候，会将对应的码值做一个输出
	// 字母、数字、标点等字符，底层逻辑是按照ASCII进行存储

	var c4 rune = '中'
	fmt.Println(c4)
	fmt.Printf("c4 is %c \n", c4)
	// 汉字字符底层对应的是Unicode码值，ASCII是Unicode的钱128位
	// go使用的是UTF-8编码，Unicode是对应的字符集，UTF-8是Unicode的其中一种编码方案

	// 转义字符
	// \n 换行
	fmt.Println("aaa\nbbb")
	// \b 退格
	fmt.Println("aaa\bccc")
	// \t 制表符 一个制表符 = 8位
	fmt.Println("aaaaaaaaaaaaaaaa")
	fmt.Println("aaaaa\taaaaaaaaaaa")
	fmt.Println("aaaa\taaaaaaaaaaa")
	fmt.Println("aaaaaaaa\taaaaaaaaaaa")

	// 字符串类型
	//	string中每个元素叫“字符”，字符有两种
	//	byte: 1个字节，代表ASII码的一个字符
	//	rune：4个字节，代表一个UTF-8字符串，一个汉字可用一个rune表示
	//  string底层是byte数组，string的长度就是该byte的数组长度，UTF-8编码下一个汉字占3个byte，即一个汉字占3个长度
	//	string可以转换为[]byte或rune类型
	//	string是常量，不能修改其中的字符

	str := "楠"
	for _, ele := range str {
		fmt.Printf("str loop %c \n", ele)
	}

	// 转byte UTF-8编码下一个汉字占3个byte
	arrByte := []byte(str)
	fmt.Printf("arrByte is %#v \n", arrByte)
	for _, ele := range arrByte {
		fmt.Printf("%d \n", ele)
	}
	fmt.Printf("arrByte length is %d, str length is %d \n", len(arrByte), len(str))

	// 转rune
	arrRune := []rune(str)
	fmt.Printf("arrRune is %#v \n", arrRune)
	for _, ele := range arrRune {
		fmt.Printf("%d \n", ele)
	}
	fmt.Printf("arrRune length is %d, str length is %d \n", len(arrRune), len(str))
}
