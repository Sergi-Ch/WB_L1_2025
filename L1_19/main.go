package main

import "fmt"

func main() {
	var str string
	var NewStr string
	_, _ = fmt.Scan(&str)
	runStr := []rune(str)
	NewStr = ""
	for i := len(runStr) - 1; i >= 0; i-- {
		NewStr += string(runStr[i])
	}
	fmt.Println(NewStr)
}
