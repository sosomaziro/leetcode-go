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
		ans = GetMax(GetMin(height[l], height[r]) * (r - l), ans)
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
			area = GetMax(GetMin(height[j], height[j+i]) * i, area)
		}
	}
	return area
}

