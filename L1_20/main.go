package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	str = reader.Text()
	tempStr := ""
	result := ""

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if tempStr != "" {
				if result != "" {
					result = tempStr + " " + result
				} else {
					result = tempStr
				}
				tempStr = ""
			}
		} else {
			tempStr += string(str[i])
		}
	}

	if tempStr != "" {
		if result != "" {
			result = tempStr + " " + result
		} else {
			result = tempStr
		}
	}

	fmt.Println(result)
}
