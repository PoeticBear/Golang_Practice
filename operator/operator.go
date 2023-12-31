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
	var num1 int = 10
	var num2 int = 3
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)

	// 逻辑运算符
	var num1 bool = true
	var num2 bool = false
	fmt.Println(num1 && num2)
	fmt.Println(num1 || num2)
	fmt.Println(!num1)

	// 位运算符
	var num1 int8 = 10
	var num2 int8 = 3
	fmt.Println(num1 & num2)
	fmt.Println(num1 | num2)
	fmt.Println(num1 ^ num2)
	fmt.Println(num1 << 2)
	fmt.Println(num1 >> 2)

	// 赋值运算符
	var num1 int = 10
	var num2 int = 3
	num1 += num2
	fmt.Println(num1)
	num1 -= num2
	fmt.Println(num1)
	num1 *= num2
	fmt.Println(num1)
	num1 /= num2
	fmt.Println(num1)
	num1 %= num2
	fmt.Println(num1)
	num1 &= num2
	fmt.Println(num1)
	num1 |= num2
	fmt.Println(num1)
	num1 ^= num2
	fmt.Println(num1)
	num1 <<= num2
	fmt.Println(num1)
	num1 >>= num2
	fmt.Println(num1)

}
