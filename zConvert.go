package main

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
*/
func main() {
	println(reverse(1534236469))
	println(convert("abcdef", 2))
}

func reverse(x int) int {
	if x > 2147483647 || x < -2147483648 {return 0}
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > 2147483647 || result < -2147483648 {return 0}
	return result
}

func convert(s string, numRows int) string {
	result := ""
	if numRows == 1 {
		result = s
		return result
	}
	rs := make([]string, numRows)
	r := 0
	goingDown := false
	for i := 0; i < len(s); i++ {
		rs[r] += string(s[i])
		if r == 0 || r == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			r++
		} else {
			r--
		}
	}
	for _, sub := range rs {
		result += sub
	}
	return result
}
