package main

import "strings"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
*/
func main() {
	println(longestPalindrome("ccc"))
}

func longestPalindrome(s string) string {
	result := ""
	str := ""
	li := 0
	ri := 0
	for li < len(s) {
		for ri < len(s) {
			i := strings.Index(str, string(s[ri]))
			str += string(s[ri])
			if i != -1 {
				str = str[i:]
				result = getLongerString(result, str)
				ri++
				break
			}
			if len(result) < 1 {
				result = str
			}
			ri++
		}
		li++
	}
	return result
}

func getLongerString(a, b string) string {
	if len(a) >= len(b) {
		return a
	}
	return b
}
