package main

import "strings"

var justString string

func createHugeString(count int) string {
	var stri string
	stri = "1"
	stri = strings.Repeat(stri, count)
	return stri
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
}
