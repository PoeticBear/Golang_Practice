package flow

func RunExample() {
	// 条件判断
	var num1 int = 10
	var num2 int = 3
	if num1 > num2 {
		println("num1大于num2")
	} else {
		println("num1小于num2")
	}

	// switch语句
	var num3 int = 10
	switch num3 {
	case 1:
		println("num3等于1")
	case 2:
		println("num3等于2")
	case 3:
		println("num3等于3")
	default:
		println("num3不等于1、2、3")
	}

	// for循环
	for i := 0; i < 10; i++ {
		println(i)
	}

}
