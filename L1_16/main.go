package main

import "fmt"

func quickSort(mass []int) []int {
	if len(mass) <= 1 {
		return mass
	}
	pivot := mass[len(mass)/2]
	low := make([]int, 0)
	middle := make([]int, 0)
	high := make([]int, 0)
	for _, val := range mass {
		if val < pivot {
			low = append(low, val)
		}
		if val == pivot {
			middle = append(middle, val)
		}
		if val > pivot {
			high = append(high, val)
		}
	}

	sortedLow := quickSort(low)
	sortedHigh := quickSort(high)
	result := append(append(sortedLow, middle...), sortedHigh...)
	return result
}

func main() {
	test := []int{3, 43, 2, 5, 2, 3, 6, 32, 5}
	fmt.Println(quickSort(test))
}
