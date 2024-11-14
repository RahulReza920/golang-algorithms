package main

import "fmt"

func isEven(val int) bool {

	switch {
	case val%2 == 0:
		fmt.Println("it is even")
		return true
	default:
		fmt.Println("it is odd")

		return false
	}
}

func limitCheck(first, second int) bool {

	switch value := 100; {
	case value > first && value < second:
		fmt.Println("value in range")
		return true
	case value < first || value > second:
		fmt.Println("value not in range")
		return false

	}
	return true
}

// func main() {

// 	num := 3

// 	isEven(num)
// 	limitCheck(10,50)
// 	limitCheck(10,500)
// }
