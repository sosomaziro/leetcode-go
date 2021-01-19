package main

func main() {
	strs := []string{"dd","df", "ddvc","c"}

	println(longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var getCommonPrefix func(start, end int) string
	getCommonPrefix = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		leftStr := getCommonPrefix(start, mid)
		rightStr := getCommonPrefix(mid + 1, end)
		minLen := GetMin(len(leftStr), len(rightStr))
		for i:= 0; i<minLen;i++ {
			if leftStr[i] != rightStr[i] {
				return leftStr[:i]
			}
		}
		return leftStr[:minLen]
	}
	return getCommonPrefix(0, len(strs)-1)
}