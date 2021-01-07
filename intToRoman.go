package main

func main() {
	println(intToRoman(9))
	println(intToRoman(4))
	println(intToRoman(90))
	println(romanToInt("MCMXCIV"))
	println(romanToInt("LVIII"))
	println(romanToInt("IX"))
}

func romanToInt(s string) int {
	result := 0
	getRes := func(n int, c string) {
		for len(s) >= len(c) {
			ch := s[0:len(c)]
			if ch == c {
				result += n
				s = s[len(c):]
			} else {
				break
			}
		}
	}
	getRes(1000,"M")
	getRes(900,"CM")
	getRes(500,"D")
	getRes(400,"CD")
	getRes(100,"C")
	getRes(90,"XC")
	getRes(50,"L")
	getRes(40,"XL")
	getRes(10,"X")
	getRes(9,"IX")
	getRes(5,"V")
	getRes(4,"IV")
	getRes(1,"I")
	return result
}

func intToRoman(num int) string {
	result := ""
	getRes := func(n int, c string) {
		for i:=0;i < num / n;i++ {
			result += c
		}
		num %= n
	}
	getRes(1000,"M")
	getRes(900,"CM")
	getRes(500,"D")
	getRes(400,"CD")
	getRes(100,"C")
	getRes(90,"XC")
	getRes(50,"L")
	getRes(40,"XL")
	getRes(10,"X")
	getRes(9,"IX")
	getRes(5,"V")
	getRes(4,"IV")
	getRes(1,"I")

	return result
}