package main

import "fmt"

func binSort(ind int, data []int) int {
	if (len(data) == 1) && (data[0] != ind) {
		return -1
	} else if (len(data) == 1) && (data[0] == ind) {
		return 0
	}

	midIndex := len(data) / 2
	midValue := data[midIndex]

	if midValue > ind {
		return binSort(ind, data[:midIndex])
	} else if midValue == ind {
		return 0
	} else if midValue < ind {
		result := binSort(ind, data[midIndex:])
		if result >= 0 {
			return midIndex + result + 1
		}
		return -1
	}
	return -2
}

func main() {

	data := []int{1, 2, 4, 5, 6, 7}
	answer := binSort(1, data)
	fmt.Println(answer)

}
