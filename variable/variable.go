package variable

import "fmt"

func RunCode() {
	var (
		name     string
		age      int
		location string
	)

	fmt.Printf("%q %d %q\n", name, age, location)

	// 初始化赋值
	var (
		name1     = "naveen"
		age1      = 29
		location1 = "Berlin"
	)

	fmt.Printf("%q %d %q\n", name1, age1, location1)

	// 短变量声明
	name2, age2 := "naveen", 29
	fmt.Printf("%q %d\n", name2, age2)

	// 批量声明
	var (
		name3     = "naveen"
		age3      = 29
		location3 = "Berlin"
	)

	fmt.Printf("%q %d %q\n", name3, age3, location3)

}
