package main

func main() {
	println(isPalindrome(0))
}

func isPalindrome(x int) bool {
	if x >= 0 {
		t := x
		r := 0
		for t != 0 {
			r = r*10 + t%10
			t /= 10
		}
		return r == x
	}
	return false
}
