package numtype

import (
	"fmt"
	"unsafe"
)

func RunExample() {
	// 有符号整型
	var num int
	fmt.Println(num)
	fmt.Printf("%d\n", num)
	fmt.Printf("%T\n", num)

	var num1 int8
	fmt.Println(num1)
	fmt.Printf("%d\n", num1)
	fmt.Printf("%T\n", num1)

	var num2 int16
	fmt.Println(num2)
	fmt.Printf("%d\n", num2)
	fmt.Printf("%T\n", num2)

	var num3 int32
	fmt.Println(num3)
	fmt.Printf("%d\n", num3)
	fmt.Printf("%T\n", num3)

	var num4 int64
	fmt.Println(num4)
	fmt.Printf("%d\n", num4)
	fmt.Printf("%T\n", num4)

	// 无符号整型
	var num5 uint
	fmt.Println(num5)
	fmt.Printf("%d\n", num5)
	fmt.Printf("%T\n", num5)

	var num6 uint8
	fmt.Println(num6)
	fmt.Printf("%d\n", num6)
	fmt.Printf("%T\n", num6)

	var num7 uint16
	fmt.Println(num7)
	fmt.Printf("%d\n", num7)
	fmt.Printf("%T\n", num7)

	var num8 uint32
	fmt.Println(num8)
	fmt.Printf("%d\n", num8)
	fmt.Printf("%T\n", num8)

	var num9 uint64
	fmt.Println(num9)
	fmt.Printf("%d\n", num9)
	fmt.Printf("%T\n", num9)

	// 浮点型（IEEE-754标准）
	var num10 float32
	fmt.Println(num10)
	fmt.Printf("%f\n", num10)
	fmt.Printf("%T\n", num10)

	var num11 float64
	fmt.Println(num11)
	fmt.Printf("%f\n", num11)
	fmt.Printf("%T\n", num11)

	var num12 complex64
	fmt.Println(num12)
	fmt.Printf("%f\n", num12)
	fmt.Printf("%T\n", num12)

	var num13 complex128
	fmt.Println(num13)
	fmt.Printf("%f\n", num13)
	fmt.Printf("%T\n", num13)

	var num14 byte
	fmt.Println(num14)
	fmt.Printf("%d\n", num14)
	fmt.Printf("%T\n", num14)

	// sizeof 函数 返回变量的内存大小
	fmt.Println(unsafe.Sizeof(num))

	// 最大值和最小值
	fmt.Println(^uint(0))
	fmt.Println(^uint8(0))

	// 十进制输出
	fmt.Printf("%d\n", 42)
	fmt.Printf("%d\n", 0600)
	fmt.Printf("%d\n", 0o600)
	fmt.Printf("%d\n", 0x600)

	// 八进制输出
	fmt.Printf("%o\n", 42)
	fmt.Printf("%o\n", 0600)
	fmt.Printf("%o\n", 0o600)
	fmt.Printf("%o\n", 0x600)

	// 十六进制输出
	fmt.Printf("%x\n", 42)
	fmt.Printf("%x\n", 0600)
	fmt.Printf("%x\n", 0o600)
	fmt.Printf("%x\n", 0x600)

}
