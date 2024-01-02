package homework

import (
	"fmt"
	"math"
)

// 奇数或偶数
func OddOrEven(num int) {
	if num%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}
}

func SumDivisibleByThree() {
	var sum int
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			sum += i
		}
	}

	fmt.Println("result is ", sum)
}

func Print99MultiTabel() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
	}
}

func CheckPrimeNumber(num int) {
	if num < 2 {
		fmt.Println("不是质数")
		return
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			fmt.Println("不是质数")
			return
		}
	}

	fmt.Println("是质数")
}

func CheckPalindrome(str string) {
	runeStr := []rune(str)
	len := len(runeStr)
	for i := 0; i < len/2; i++ {
		if runeStr[i] != runeStr[len-1-i] {
			fmt.Println("不是回文字符串")
			return
		}
	}
	fmt.Println("是回文字符串")
}

func CalFactorial(n int) {
	var result int = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("%d的阶乘是：%d\n", n, result)
}

func printFiboN(n int) {
	if n >= 1 {
		fmt.Println("0")
	}

	if n >= 2 {
		fmt.Println("1")
	}

	if n >= 3 {
		a, b := 0, 1
		for i := 3; i <= n; i++ {
			c := a + b
			fmt.Println(c)
			a, b = b, c
		}
	}
}

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func RollDice() {
	// 随机数种子

}
