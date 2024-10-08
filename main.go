package main

import (
	"fmt"
)

func main() {
	var printValue string = "Hello"
	printMe(printValue)

	var num1 int = 1
	var num2 int = 2
	var result, product int = practiceNumbers(num1, num2)

	fmt.Println(result, product)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func practiceNumbers(num1 int, num2 int) (int, int) {
	var result int = num1 + num2
	var product int = num1 * num2
	return result, product
}
