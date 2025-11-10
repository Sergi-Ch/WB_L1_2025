package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Intersection(fData []int, sData []int) []int {
	var results []int
	setMap := make(map[int]int)
	for _, val := range fData {
		_, ok := setMap[val]
		if ok {
			continue
		} else {
			setMap[val] = 1
		}

	}
	for _, val := range sData {

		if v := setMap[val]; v == 1 {
			results = append(results, val)
			setMap[val] += 1
		}
	}
	return results
}

func main() {

	var fSlice []int
	var sSlice []int
	var input []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Заполнение 1 слайса, ввести числа через пробел >>> ")
	tempData, _ := reader.ReadString('\n')
	tempData = strings.TrimSpace(tempData)
	input = strings.Split(tempData, " ")
	for _, val := range input {
		v, _ := strconv.Atoi(val)
		fSlice = append(fSlice, v)
	}
	fmt.Println("Заполнение 2 слайса, ввести числа через пробел >>> ")
	tempData, _ = reader.ReadString('\n')
	tempData = strings.TrimSpace(tempData)
	input = strings.Split(tempData, " ")
	for _, val := range input {
		v, _ := strconv.Atoi(val)
		sSlice = append(sSlice, v)
	}
	results := Intersection(fSlice, sSlice)
	fmt.Println(results)

}
