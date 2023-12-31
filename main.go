package main

import (
	"Golang_Practice/user"
	"Golang_Practice/variable"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	result := user.HelloUser()
	fmt.Println(result)

	variable.RunCode()

}
