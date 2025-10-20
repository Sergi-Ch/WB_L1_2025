package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	groups := make(map[int][]float64)
	var data []float64
	var input string

	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	tempData := strings.Split(input, ",")

	for _, val := range tempData {
		val = strings.TrimSpace(val)
		flVal, _ := strconv.ParseFloat(val, 64)
		data = append(data, flVal)
	}

	for _, val := range data {
		var group int
		group = int(val/10) * 10

		groups[group] = append(groups[group], val)
	}
	fmt.Println(groups)
}
