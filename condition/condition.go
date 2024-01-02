package condition

func RunExample() {
	// if语句
	var num1 int = 10
	var num2 int = 3
	if num1 > num2 {
		println("num1大于num2")
	}

	// if...else语句
	var num3 int = 10
	var num4 int = 3
	if num3 > num4 {
		println("num3大于num4")
	} else {
		println("num3小于num4")
	}

	// if...else if...else语句
	var num5 int = 10
	var num6 int = 3
	if num5 > num6 {
		println("num5大于num6")
	} else if num5 < num6 {
		println("num5小于num6")
	} else {
		println("num5等于num6")
	}

	// if嵌套语句
	var num7 int = 10
	var num8 int = 3
	if num7 > num8 {
		println("num7大于num8")
		if num7 > 5 {
			println("num7大于5")
		}
	} else {
		println("num7小于num8")
	}

	// switch语句
	var num9 int = 10
	switch num9 {
	case 1:
		println("num9等于1")
	case 2:
		println("num9等于2")
	case 3:
		println("num9等于3")
	default:
		println("num9不等于1、2、3")
	}

	// switch语句
	var num10 int = 10
	switch num10 {
	case 1, 2, 3:
		println("num10等于1、2、3")
	default:
		println("num10不等于1、2、3")
	}

	// switch语句
	var num11 int = 10
	switch {
	case num11 > 0 && num11 < 5:
		println("num11大于0小于5")
	case num11 > 5 && num11 < 10:
		println("num11大于5小于10")
	default:
		println("num11不在0~10之间")
	}

}
