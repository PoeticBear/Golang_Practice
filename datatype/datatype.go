package datatype

import "fmt"

func RunExample() {
	var name string
	var age int
	var isLogin bool
	var weight float32
	var height float64
	var a byte
	var b rune
	var c complex64
	var d complex128

	// 赋值
	name = "naveen"
	age = 29
	isLogin = true
	weight = 60.5
	height = 172.5
	a = 'a'
	b = 'b'
	c = 1 + 2i
	d = 1 + 2i

	// 打印
	fmt.Println(name, age, isLogin, weight, height, a, b, c, d)

	// 格式化打印
	fmt.Printf("%s %d %t %f %f %c %c %v %v\n", name, age, isLogin, weight, height, a, b, c, d)

	// 打印类型
	fmt.Printf("%T %T %T %T %T %T %T %T %T\n", name, age, isLogin, weight, height, a, b, c, d)

	// 指针类型
	var name1 *string
	var age1 *int
	var isLogin1 *bool
	var weight1 *float32
	var height1 *float64
	var a1 *byte
	var b1 *rune
	var c1 *complex64
	var d1 *complex128

	// 赋值
	name1 = &name
	age1 = &age
	isLogin1 = &isLogin
	weight1 = &weight
	height1 = &height
	a1 = &a
	b1 = &b
	c1 = &c
	d1 = &d

	// 打印
	fmt.Println(name1, age1, isLogin1, weight1, height1, a1, b1, c1, d1)

	// 格式化打印
	fmt.Printf("%p %p %p %p %p %p %p %p %p\n", name1, age1, isLogin1, weight1, height1, a1, b1, c1, d1)

	// 打印类型
	fmt.Printf("%T %T %T %T %T %T %T %T %T\n", name1, age1, isLogin1, weight1, height1, a1, b1, c1, d1)

	// 指针取值
	fmt.Println(*name1, *age1, *isLogin1, *weight1, *height1, *a1, *b1, *c1, *d1)

	// 复合类型
	var name2 [10]string
	var age2 [10]int
	var isLogin2 [10]bool
	var weight2 [10]float32
	var height2 [10]float64
	var a2 [10]byte
	var b2 [10]rune
	var c2 [10]complex64
	var d2 [10]complex128

	// 赋值
	name2[0] = "naveen"
	age2[0] = 29
	isLogin2[0] = true
	weight2[0] = 60.5
	height2[0] = 172.5
	a2[0] = 'a'
	b2[0] = 'b'
	c2[0] = 1 + 2i
	d2[0] = 1 + 2i

	// 打印
	fmt.Println(name2, age2, isLogin2, weight2, height2, a2, b2, c2, d2)

	// 格式化打印
	fmt.Printf("%s %d %t %f %f %c %c %v %v\n", name2, age2, isLogin2, weight2, height2, a2, b2, c2, d2)

	// 函数类型
	var name3 func() string
	var age3 func() int
	var isLogin3 func() bool
	var weight3 func() float32
	var height3 func() float64
	var a3 func() byte
	var b3 func() rune
	var c3 func() complex64
	var d3 func() complex128

	// 赋值
	name3 = getName
	age3 = getAge
	isLogin3 = getIsLogin
	weight3 = getWeight
	height3 = getHeight
	a3 = getA
	b3 = getB
	c3 = getC
	d3 = getD

	// 打印
	fmt.Println("函数类型")
	fmt.Println(name3, age3, isLogin3, weight3, height3, a3, b3, c3, d3)

	// 格式化打印
	fmt.Printf("%s %d %t %f %f %c %c %v %v\n", name3, age3, isLogin3, weight3, height3, a3, b3, c3, d3)

}

// 定义上面的函数
func getName() string {
	return "naveen"
}

func getAge() int {
	return 29
}

func getIsLogin() bool {
	return true
}

func getWeight() float32 {
	return 60.5
}

func getHeight() float64 {
	return 172.5
}

func getA() byte {
	return 'a'
}

func getB() rune {
	return 'b'
}

func getC() complex64 {
	return 1 + 2i
}

func getD() complex128 {
	return 1 + 2i
}
