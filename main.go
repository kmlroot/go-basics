package main

import (
	"fmt"

	"github.com/fmauricios/go-basics/flow"
	"github.com/fmauricios/go-basics/name"
	"github.com/fmauricios/go-basics/numbers"
)

const helloWorld string = "Hi %s %s, welcome to the awesome Golang world!\n"
const testConst = "Test"

func main() {
	lastname := "Serna"

	firstName := name.GetName()

	var number = 100

	a, b, c := numbers.GetVariables()
	f32, f64 := numbers.GetFloat()

	add := numbers.Sum(a, b)

	fmt.Printf(helloWorld, firstName, lastname)
	fmt.Println(number, a, b, c, lastname)
	fmt.Println(add)
	fmt.Println(f32, f64)

	name.String2()
	flow.SwitchTest()
}
