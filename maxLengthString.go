package main

/*
无重复字符的最长子串
*/
func main() {
	println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) (result int) {
	i1 := 0
	i2 := 0
	cs := map[byte]int{}
	l := len(s)
	for i1 < l {
		if i1 != 0 {
			delete(cs, s[i1-1])
		}

		for i2 < l && cs[s[i2]] == 0 {
			cs[s[i2]] = 1
			i2++
			result = maxInt(result, len(cs))
		}
		i1++
	}
	return
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
