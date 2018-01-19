package main

import "fmt"

func main() {
	number1 := 0
	number2 := 0

	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &number1)

	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d", &number2)

	result1 := add(number1, number2)
	result2 := substract(number1, number2)

	fmt.Printf("Add: %d", result1)
	fmt.Println("")
	fmt.Printf("Substract: %d", result2)
}

func add(a int, b int) int {
	return a + b
}

func substract(a int, b int) int {
	return a - b
}
