package main

import (
	"fmt"
	"time"

	"github.com/fmauricios/go-basics/flow"
	"github.com/fmauricios/go-basics/maps"
	"github.com/fmauricios/go-basics/name"
	"github.com/fmauricios/go-basics/numbers"
	"github.com/fmauricios/go-basics/structs"
)

const helloWorld string = "Hi %s %s, welcome to the awesome Golang world!\n"
const testConst = "Test"

func main() {
	lastname := "Serna"

	firstName := name.GetName()

	var number = 100

	a, b, c := numbers.GetVariables()
	f32, f64 := numbers.GetFloat()

	// add := numbers.Sum(a, b)

	fmt.Printf(helloWorld, firstName, lastname)
	fmt.Println(number, a, b, c, lastname)
	// fmt.Println(add)
	fmt.Println(f32, f64)

	name.String2()
	flow.SwitchTest()

	fmt.Println(maps.GetMap("Mauricio"))

	structs.InterfaceTest()

	number2, err := numbers.Sum(50, 50)

	if err != nil {
		panic(err)
	}

	fmt.Println(number2)

	pointerTest()

	go forGo(500)
	go forGo(400)
	time.Sleep(10000 * time.Millisecond)
}

func pointerTest() {
	a := 100
	var b *int
	b = &a

	fmt.Println("Without modifitication")

	fmt.Println(a, *b)
	fmt.Println(&a, b)

	pointerModify(b)
}

func helloGo(index int) {
	fmt.Println("Println in helloGo Go rutine #", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}

func pointerModify(c *int) {
	*c = 10
	fmt.Println("With modification")
}
