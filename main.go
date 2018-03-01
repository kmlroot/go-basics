package main

import (
	"fmt"
	"strings"

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

	string2()
	flow.SwitchTest()
}

func string2() {
	var text = "Mauricio Serna, Alba Florez, Hello Golang"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Alba", "Luz Albita", -1))
	fmt.Println(strings.Split(text, ","))
}
