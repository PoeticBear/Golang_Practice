package formattedoutput

import "fmt"

type WebSite struct {
	Name string
	Url  string
}

func RunExample() {
	site := WebSite{"C语言中文网", "http://c.biancheng.net/golang/"}
	fmt.Printf("%v\n", site)  // {C语言中文网 http://c.biancheng.net/golang/}
	fmt.Printf("%+v\n", site) // {Name:C语言中文网 Url:http://c.biancheng.net/golang/}
	fmt.Printf("%#v\n", site) // main.WebSite{Name:"C语言中文网", Url:"http://c.biancheng.net/golang/"}
	fmt.Printf("%T\n", site)  // main.WebSite
	fmt.Printf("%%\n")        // % 百分号标记 %v %T 和 %% 之间的区别
	fmt.Printf("%t\n", true)  // true
	fmt.Printf("%b\n", 1024)  // 10000000000
	fmt.Printf("%c\n", 111)   // o
	fmt.Printf("%d\n", 10)    // 10
	fmt.Printf("%o\n", 8)     // 10
	fmt.Printf("%q\n", 22)    // "22"
	fmt.Printf("%x\n", 1223)  // 4c7
	fmt.Printf("%X\n", 1223)  // 4C7
	fmt.Printf("%U\n", 123)   // U+007B
	fmt.Printf("%e\n", 123.4) // 1.234000e+02
	fmt.Printf("%E\n", 123.4) // 1.234000E+02
	fmt.Printf("%f\n", 123.4) // 123.400000
	fmt.Printf("%g\n", 123.4) // 123.4
	fmt.Printf("%G\n", 123.4) // 123.4
	fmt.Printf("%s\n", "abc") // abc
	fmt.Printf("%q\n", "abc") // "abc"
	fmt.Printf("%p\n", &site) // 0xc00000c080
}
