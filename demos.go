package main

import "reflect"

func main() {
	stringTest()
}

func mapDemos() {
	m := map[int]string{}
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	v, ok := m[4]
	if ok {
		println(v)
	} else {
		println(len(v))
	}
}

func stringTest() {
	str := "abcdefG中"
	for i, i2 := range str {
		println(i, " -> ", i2, " -> ", reflect.TypeOf(i2))
	}
}
