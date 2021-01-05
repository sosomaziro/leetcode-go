package main

import "fmt"

func main() {
	//println(isMatch("aaa", "ab*a*c*a"))
	println(isMatch("b", "a*b"))
}

//func isMatch(s string, p string) bool {
//	ic := 0
//	ip := 0
//	lpc := ""
//	var sc string
//	for {
//		if ic >= len(s) {
//			if ip >= len(p) {
//				return true
//			}
//			sc = "aa"
//		} else {
//			sc = string(s[ic])
//		}
//
//		if ip >= len(p) {
//			return false
//		}
//
//		pc := string(p[ip])
//
//		if pc == "." {
//			ic++
//			ip++
//			lpc = pc
//		} else if pc == sc {
//			ic++
//			ip++
//			lpc = pc
//		} else if pc == "*" {
//			if sc == lpc || lpc == "." {
//				ic++
//			} else {
//				ip++
//				lpc = ""
//			}
//		} else {
//			if ip+1 >= len(p) || string(p[ip+1]) != "*" {
//				return false
//			}
//			ip++
//		}
//	}
//}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i-1][j-1]
			}
		}
	}
	for _, bools := range f {
		fmt.Println(bools)
	}
	return f[m][n]
}
