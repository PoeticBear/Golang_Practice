package operator

import "fmt"

func RunExample() {
	// 算术运算符
	var num1 int = 10
	var num2 int = 3
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	// 关系运算符
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 > num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 <= num2)

	// 逻辑运算符
	var isLogin bool = true
	var isLogout bool = false
	fmt.Println(isLogin && isLogout)
	fmt.Println(isLogin || isLogout)
	fmt.Println(!isLogin)

	// 位运算符
	var num3 int = 60 // 0011 1100
	var num4 int = 13 // 0000 1101
	fmt.Println(num3 & num4)
	fmt.Println(num3 | num4)
	fmt.Println(num3 ^ num4)
	fmt.Println(num3 << 2)
	fmt.Println(num3 >> 2)

	// 赋值运算符
	var num5 int = 10
	num5 += 2
	fmt.Println(num5)
	num5 -= 2
	fmt.Println(num5)
	num5 *= 2
	fmt.Println(num5)
	num5 /= 2
	
	
}
