package variable

import "fmt"

func UseConst() {
	// 声明常量
	const sex = "男"
	println(sex)
	// 批量声明
	const (
		PI  = 3.14
		PI1 = 3.15
	)

	fmt.Printf("PI: %f PI1: %f \n", PI, PI1) //
	// 0 1
	const (
		a0 = iota // 0
		a1 = iota // 1
	)

	fmt.Printf("a0: %d a1: %d \n", a0, a1) // 0 1

	const (
		PI2 = 100
		PI3 // 100跟上一行的值相同
		PI4 // 100
	)
	fmt.Printf("PI2: %d PI3: %d PI4: %d \n", PI2, PI2, PI4)
	// iota
	const (
		_  = iota      // 默认是0, 忽略0
		v1             // 1
		v2             // 2
		_              // 3 忽略这个值
		v4             // 4
		v5 = iota + 10 // 15
	)

	fmt.Printf("v1: %d v2: %d v4: %d v5: %d \n", v1, v2, v4, v5)

	const (
		g0 = iota // 默认是0
		g1 = 30   // 1
		g2 = iota // 2
		g3        // 3
	)
	fmt.Printf("g0: %d g1: %d g2: %d g3: %d \n", g0, g1, g2, g3)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
	)

	fmt.Printf("KB: %d MB: %d GB: %d TB: %d \n", KB, MB, GB, TB)

	const (
		h1, h2 = iota + 1, iota + 2 // iota = 0
		h3, h4
	)
	fmt.Printf("h1: %d h2: %d h3: %d h4: %d \n", h1, h2, h3, h4)
}
