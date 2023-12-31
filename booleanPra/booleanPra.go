package booleanPra

import "fmt"

func RunExample() {
	var isLogin bool

	// 赋值
	isLogin = true

	// 打印
	fmt.Println(isLogin)

	// 格式化打印
	fmt.Printf("%t\n", isLogin)

	// 打印类型
	fmt.Printf("%T\n", isLogin)

	// 指针类型
	var isLogin1 *bool

	// 赋值
	isLogin1 = &isLogin

	// 打印
	fmt.Println(isLogin1)

	// 打印类型
	fmt.Printf("%T\n", isLogin1)

	// 0和1不能直接转换为bool类型

}
