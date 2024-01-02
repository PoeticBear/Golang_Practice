package homework04

import "fmt"

func add(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func sum(nums ...int) int {
	numSum := 0
	for _, num := range nums {
		numSum += num
	}
	return numSum
}

// 匿名函数
func anonymousFunc() {
	func() {
		fmt.Println("匿名函数")
	}()

	greet := func() {
		fmt.Println("匿名函数")
	}
	greet()
}

// 闭包
func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
