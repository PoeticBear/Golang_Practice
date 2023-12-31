package constant

import "fmt"

func RunExample() {
	// 常量声明, 不能使用 := 语法, 必须使用 = 语法, 且必须初始化, 不能只声明不赋值
	// 赋值后不能修改

	const pi = 3.1415
	// pi = 3 // cannot assign to pi
	fmt.Println(pi)

	// 批量声明
	const (
		statusOk = 200
		notFound = 404
	)

	// 类型推断
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Printf("%T %v\n", n1, n1)

	// iota

	const (
		a1 = iota // 0
		a2        // 1
		a3        // 2
	)
	fmt.Println(a1, a2, a3)

	// 跳值
	const (
		b1 = iota // 0
		b2 = 100  // 100
		b3 = iota // 2
		b4        // 3
	)
	fmt.Println(b1, b2, b3, b4)

	// 定义数量级
	const (
		_  = iota
		KB = 1 << (10 * iota) // 1 << (10 * 1)
		MB                    // 1 << (10 * 2)
		GB                    // 1 << (10 * 3)
		TB                    // 1 << (10 * 4)
		PB                    // 1 << (10 * 5)
	)
	fmt.Println(KB, MB, GB, TB, PB)

	// 多个 iota 定义在一行
	const (
		c1, c2 = iota + 1, iota + 2 // 1, 2
		c3, c4 = iota + 1, iota + 2 // 2, 3
	)
	fmt.Println(c1, c2, c3, c4)

	// 定义常量组
	const (
		d1 = iota // 0
		d2 = 100  // 100
		d3 = iota // 2
		d4        // 3
	)
	fmt.Println(d1, d2, d3, d4)

	// 更多关于 iota 的用法
	const (
		e1, e2 = iota + 1, iota + 2 // 1, 2
		_, _                        // 2, 3
		e3, e4                      // 3, 4
	)
	fmt.Println(e1, e2, e3, e4)

}
