package helper

func TotalPage(count, limit uint) uint {
	if limit == 0 {
		return 0
	}
	return count/limit + 1
}

func SkipPage(page, limit uint) uint {
	if page == 0 {
		page = 1
	}

	return (page - 1) * limit
}

func RainbowPage(current, count, display int) []int {
	e, left, right, length := display%2 == 0, display/2, display/2, display
	if e {
		right++
	}
	if count < display {
		length = count
	}

	result := make([]int, length)
	if count >= display {
		switch {
		case current <= left:
			for i := 0; i < length; i++ {
				result[i] = i + 1
			}
		case current > count-right:
			for i := 0; i < length; i++ {
				result[i] = i + count - display + 1
			}
		default:
			for i := 0; i < length; i++ {
				result[i] = i + current - left + booleanToInt(e)
			}
		}

		return result
	}
	for i := 0; i < length; i++ {
		result[i] = i + 1
	}

	return result
}

func booleanToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

