package utils

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
