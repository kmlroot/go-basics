package structs

import (
	"fmt"
)

// It's a interface of Course and Career
type MyInterface interface {
	Subscribe(name string)
}

func deferTest() {
	fmt.Println("The function interface has finished")
}

func InterfaceTest() {
	defer deferTest()
	course := Course{Name: "Golang", Slug: "golang", Skills: []string{"1", "2"}}
	career := new(Career)
	career.Name = "Golang"
	career.Slug = "go"

	callSubscribe(course)
	callSubscribe(career)
}

func callSubscribe(m MyInterface) {
	m.Subscribe("Mauricio")
}
