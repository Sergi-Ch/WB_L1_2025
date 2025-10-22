package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateSet(fData []string) []string {
	var results []string
	setMap := make(map[string]int)

	for _, val := range fData {
		val = strings.TrimSpace(val)
		_, ok := setMap[val]
		if ok {
			continue
		} else {
			setMap[val] = 1
			results = append(results, val)
		}

	}

	return results
}

func main() {
	var fSlice []string
	var input []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Заполнение слайса, ввести слова через запятую >>> ")
	tempData, _ := reader.ReadString('\n')
	tempData = strings.TrimSpace(tempData)
	input = strings.Split(tempData, ",")
	fSlice = CreateSet(input)
	fmt.Println(fSlice)
}
