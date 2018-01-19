package main

import "fmt"

const helloWorld string = "Hi %s %s, welcome to the awesome Golang world!\n"
const testConst = "Test"

func main() {
	var name string
	lastname := "Serna"

	var number = 100

	var (
		a = 1
		b = 2
		c = 3
	)

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf(helloWorld, name, lastname)
	fmt.Println(number, a, b, c, lastname)
}
