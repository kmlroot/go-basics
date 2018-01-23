package main

import "fmt"

const helloWorld string = "Hi %s %s, welcome to the awesome Golang world!\n"
const testConst = "Test"

func main() {
	name := getName()
	lastname := "Serna"

	var number = 100

	a, b, c := getVariables()
	f32, f64 := getFloat()

	add := sum(a, b)

	fmt.Printf(helloWorld, name, lastname)
	fmt.Println(number, a, b, c, lastname)
	fmt.Println(add)
	fmt.Println(f32, f64)
}

func getName() string {
	var name string
	name = "Without name"

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)

	return name
}

func getVariables() (int, int, int) {
	return 1, 2, 3
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func sum(a int, b int) int {
	return a + b
}
