package flow

import "fmt"

func SwitchTest() {
	var number = 0
	fmt.Print("[Switch] Enter a number from 1 to 10: ")
	fmt.Scanf("%d", &number)

	switch number {
	case 1:
		fmt.Println("The number is 1")
	default:
		fmt.Println("The number isn't 1")
	}

	switch {
	case number%2 == 0:
		fmt.Println("The number is even")
	default:
		fmt.Println("The number is odd")
	}
}
