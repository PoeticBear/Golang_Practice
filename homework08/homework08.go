package homework08

import (
	"fmt"
	"os"
)

func ExcuteTest() {
	createDir()
}

func createDir() {
	err := os.Mkdir("new_directory", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("创建目录成功")
}

func checkDir() {

}
