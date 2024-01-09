package osFile

import (
	"fmt"
	"io"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func createDir() {
	err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 创建多级目录
func createDirAll() {
	err := os.MkdirAll("test/a/b", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除
func remove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除目录
func removeDir() {
	err := os.RemoveAll("test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 获取工作空间
func wd() {
	dir, _ := os.Getwd()

	fmt.Printf("dir: %v\n", dir)
	os.Chdir("/Users/pengsihang/Developer")

	dir2, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir2)
}

// 重命名
func rename() {
	err := os.Rename("a.txt", "a_new.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 读文件
func readFile() {
	b, _ := os.ReadFile("a_new.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

// 写文件
func writeFile() {
	s := "hello pengsihang"
	os.WriteFile("test2.txt", []byte(s), os.ModePerm)
}

//// file 相关

// 打开或关闭文件
func openAndCloseFile() {
	// f, err := os.Open("test2.txt")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// } else {
	// 	fmt.Printf("f.Name(): %v\n", f.Name())
	// 	f.Close()
	// }

	f2, err2 := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err2 != nil {
		fmt.Printf("err: %v\n", err2)
	} else {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
		f2.Close()
	}
}

func createFileTemp() {
	f, _ := os.Create("a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

func readOpsreadOps() {
	f, _ := os.Open("test2.txt")

	for {
		buf := make([]byte, 10)
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
	}
	f.Close()
}

func readOps2() {
	f, _ := os.Open("test2.txt")
	buf := make([]byte, 3)
	n, _ := f.ReadAt(buf, 3)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
}

func readSeek() {
	f, err := os.Open("test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.Seek(3, 0)
	buf := make([]byte, 12)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()
}

func readDir() {
	de, _ := os.ReadDir("testDir")
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func ExcuteTest() {
	readSeek()
}
