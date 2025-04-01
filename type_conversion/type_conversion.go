package type_conversion

import (
	"fmt"
	"math"
	"strconv"
)

func TypeConversion() {
	// go在不同类型之间需要显示的转换
	// T(v) - 将值v转换为变量T
	//	byte和int可以互相转换
	//	float和int可以相互转换
	//	bool和int不能相互转换
	//	不同长度的int和float之间可以相互转换
	//	string可以转换为[]byte和[]rune,byte和rune可以转换为string
	//	低精度向高精度转换没问题，高精度向低精度转换会丢失位数
	//	无符号向有符号转换，最高位变为符号位
	var f float32 = 3.14
	i := int(f)
	// f还是float32类型的，这是知识将f的值3。14转换为了int类型，赋值给i
	fmt.Printf("%d \n", i)
	var s = math.Pow(2, 0)
	fmt.Printf("%f /n", s)
	var i1 int64 = 18

	// strconv 将其他类型转为string类型
	var s1 = strconv.FormatInt(i1, 10)
	fmt.Printf("s1 对应的类型是 %T, s1 = %q \n", s1, s1)
	var s2 = "true"
	// string转为bool
	var b, _ = strconv.ParseBool(s2)
	fmt.Printf("b的类型是：%T, b=%v \n", b, b)
}
