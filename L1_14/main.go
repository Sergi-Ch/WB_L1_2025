package main

import (
	"fmt"
	"reflect"
)

func CheckType(v interface{}) string {

	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"

	case float64:
		return "float64"

	case bool:
		return "bool"

	case nil:
		return "nil"

	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan { //Kind возвращает базовый тип, но он все равно из констант reflect, поэтому сравнение корректное
			return "chan"
		}
		return "unknown"
	}

}

func main() {
	ch := make(chan string)
	fmt.Println(CheckType(6))
	fmt.Println(CheckType(6.5))
	fmt.Println(CheckType("hi"))
	fmt.Println(CheckType(true))
	fmt.Println(CheckType(nil))
	fmt.Println(CheckType(ch))

}
