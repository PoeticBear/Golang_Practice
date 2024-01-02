package homework03

import (
	"fmt"
	"sort"
)

func InitArray() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}

func Print3Item() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[2])
}

func InitFruitSlice() {
	var fruits = []string{"apple", "banana", "orange"}
	fmt.Println(fruits)
}

func NewSlice() {
	var allFruits = []string{"apple", "banana", "orange", "grape", "mango"}
	var fruits = allFruits[0:3]
	fmt.Println(fruits)
}

func IterateSlice() {
	var allFruits = []string{"apple", "banana", "orange", "grape", "mango"}
	// for i := 0; i < len(allFruits); i++ {
	// 	fmt.Println(allFruits[i])
	// }
	for index, value := range allFruits {
		fmt.Println(index, value)
	}
}

func AddElement() {
	var allFruits = []string{"apple", "banana", "orange", "grape", "mango"}
	allFruits = append(allFruits, "pear")
	fmt.Println(allFruits)
}

func DeleteSecondElement() {
	var allFruits = []string{"apple", "banana", "orange", "grape", "mango"}
	allFruits = append(allFruits[:1], allFruits[2:]...)
	fmt.Println(allFruits)
}

func InitMap() {
	var fruits = map[string]int{"apple": 5, "banana": 2, "orange": 8}
	fmt.Println(fruits)

	var fruits2 = make(map[string]int)
	fruits2["apple"] = 5
	fruits2["banana"] = 2
	fruits2["orange"] = 8
	fmt.Println(fruits2)
}

func IterateMap() {
	var fruits = map[string]int{"apple": 5, "banana": 2, "orange": 8}
	for key, value := range fruits {
		fmt.Println(key, value)
	}
}

func IterateMapSorted() {
	var fruits = make(map[string]int)
	fruits["apple"] = 5
	fruits["banana"] = 2
	fruits["orange"] = 8

	var keys []string = make([]string, len(fruits))
	for key := range fruits {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, fruits[key])
	}
}

func CheckElementExist() {
	var fruits = map[string]int{"apple": 5, "banana": 2, "orange": 8}
	var value, ok = fruits["apple"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}
}
