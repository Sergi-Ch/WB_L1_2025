package main

import "fmt"

func main() {
	var number int64
	var i int64
	var endByte int64
	_, _ = fmt.Scan(&number)
	fmt.Println("введите бит, который хотите поменять (нумерация с 1 справа налево)")
	_, _ = fmt.Scan(&i)
	fmt.Println("введите бит, на который хотите поменять")
	_, _ = fmt.Scan(&endByte)
	i -= 1
	fmt.Printf("start var %d = %b \n", number, number)
	switch endByte {
	case 1:
		number |= 1 << i
	case 0:
		number &^= 1 << i

	}

	fmt.Printf("end var %d = %b \n", number, number)
}
