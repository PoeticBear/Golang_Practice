package stringPra

import (
	"bytes"
	"fmt"
	"strings"
	"unsafe"
)

func RunExample() {
	// 声明字符串
	var str string
	str = "Hello, World!"
	fmt.Println(str)

	// 格式化打印
	fmt.Printf("%s\n", str)

	// 多行字符串
	var str1 string
	str1 = `
		Hello, World!
		Hello, Golang!
	`
	fmt.Println(str1)

	// 字符串连接
	var str2 string
	str2 = "Hello, " + "World!"
	fmt.Println(str2)

	// 字符串长度
	var str3 string
	str3 = "Hello, World!"
	fmt.Println(len(str3))
	fmt.Println(unsafe.Sizeof(str3))

	// buffer.WriteString()
	var buffer bytes.Buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString("World!")
	fmt.Println(buffer.String())

	// 转义字符
	var str4 string
	str4 = "Hello, \nWorld!"
	fmt.Println(str4)

	// 原生字符串
	var str5 string
	str5 = `Hello, \nWorld!`
	fmt.Println(str5)

	// 字符串遍历
	var str6 string
	str6 = "Hello, World!"
	for i := 0; i < len(str6); i++ {
		fmt.Printf("%c\n", str6[i])
	}

	// 其他转义字符
	// \a 警告或者铃声
	// \b 退格
	// \f 换页
	// \r 回车
	// \t 制表符
	// \v 垂直制表符
	// \' 单引号
	// \" 双引号
	// \\ 反斜杠

	// 字符串切片
	var str7 string
	str7 = "Hello, World!"
	fmt.Println(str7[0:5])
	fmt.Println(str7[7:12])
	fmt.Println(str7[:5])
	fmt.Println(str7[7:])
	fmt.Println(str7[:])

	// split
	var str8 string
	str8 = "Hello, World!"
	strings.Split(str8, ",")
	fmt.Println(strings.Split(str8, ","))

	// contains
	var str9 string
	str9 = "Hello, World!"
	fmt.Println(strings.Contains(str9, "Hello"))
	fmt.Println(strings.Contains(str9, "hello"))

	// join
	var str10 []string
	str10 = []string{"Hello", "World!"}
	fmt.Println(strings.Join(str10, ", "))
	fmt.Println(strings.Join(str10, " "))

	// index
	var str11 string
	str11 = "Hello, World!"
	fmt.Println(strings.Index(str11, "Hello"))
	fmt.Println(strings.Index(str11, "hello"))

	// last index
	var str12 string
	str12 = "Hello, World!"
	fmt.Println(strings.LastIndex(str12, "Hello"))
	fmt.Println(strings.LastIndex(str12, "hello"))

	// replace
	var str13 string
	str13 = "Hello, World!"
	fmt.Println(strings.Replace(str13, "Hello", "你好", 1))
	fmt.Println(strings.Replace(str13, "Hello", "你好", 2))
	fmt.Println(strings.Replace(str13, "Hello", "你好", -1))

	// has prefix
	var str14 string
	str14 = "Hello, World!"
	fmt.Println(strings.HasPrefix(str14, "Hello"))
	fmt.Println(strings.HasPrefix(str14, "hello"))

	// has suffix
	var str15 string
	str15 = "Hello, World!"
	fmt.Println(strings.HasSuffix(str15, "World!"))
	fmt.Println(strings.HasSuffix(str15, "world!"))

	// to lower
	var str16 string
	str16 = "Hello, World!"
	fmt.Println(strings.ToLower(str16))

	// to upper
	var str17 string
	str17 = "Hello, World!"
	fmt.Println(strings.ToUpper(str17))

	// trim space
	var str18 string
	str18 = " Hello, World! "
	fmt.Println(strings.TrimSpace(str18))

	// trim left
	var str19 string
	str19 = " Hello, World! "
	fmt.Println(strings.TrimLeft(str19, " "))

	// trim right
	var str20 string
	str20 = " Hello, World! "
	fmt.Println(strings.TrimRight(str20, " "))

	// trim
	var str21 string
	str21 = " Hello, World! "
	fmt.Println(strings.Trim(str21, " "))

	// fields是按照空格分割，返回一个切片，不包含空格，如果字符串是空格，返回空切片
	var str22 string
	str22 = " Hello, World! "
	fmt.Println(strings.Fields(str22))

}
