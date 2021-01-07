package main

func main() {
	//println(isMatch("aaa", "ab*a*c*a"))
	println(maxAreaLow([]int{1, 2, 3, 4, 7, 8, 9, 10}))
	println(maxArea([]int{1, 2, 3, 4, 7, 8, 9, 10}))
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0
	for l < r {
		ans = getMax(getMin(height[l], height[r]) * (r - l), ans)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return ans
}

func maxAreaLow(height []int) int {
	area:=0
	for i:=1; i< len(height);i++{
		for j:=0 ; j < len(height) - i;j++{
			area = getMax(getMin(height[j], height[j+i]) * i, area)
		}
	}
	return area
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

