package arrayprac

import "fmt"

func RunExample() {
	// 声明
	var arr1 [5]int
	fmt.Println(arr1)

	// 赋值
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	fmt.Println(arr1)

	// 声明并赋值
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// 字符串数组
	var arr3 [5]string = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr3)

	// 字符数组
	var arr4 [5]byte = [5]byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Println(arr4)

	// 布尔数组
	var arr5 [5]bool = [5]bool{true, false, true, false, true}
	fmt.Println(arr5)

}
