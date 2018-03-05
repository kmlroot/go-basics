package main

import (
	"fmt"

	"github.com/fmauricios/go-basics/flow"
	"github.com/fmauricios/go-basics/maps"
	"github.com/fmauricios/go-basics/name"
	"github.com/fmauricios/go-basics/numbers"
)

// Course struct
type Course struct {
	Name   string
	Slug   string
	Skills []string
}

// Career struct -> Course
type Career struct {
	Course
}

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

	fmt.Println(maps.GetMap("Mauricio"))

	course := Course{Name: "Golang", Slug: "golang", Skills: []string{"1", "2"}}
	course1 := new(Course)

	course1.Name = "Golang1"
	course1.Slug = "golang1"
	course1.Skills = []string{"backend1"}

	fmt.Println(course)
}
