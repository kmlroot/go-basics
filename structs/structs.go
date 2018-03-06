package structs

import "fmt"

// Course struct
type Course struct {
	Name   string
	Slug   string
	Skills []string
}

func (c Course) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s\n", name, c.Name)
}

// Career struct -> Course
type Career struct {
	Course
}

func (c Career) Suscribe(name string) {
	fmt.Println("La persona %s se ha registrado a la carrera %s", name, c.Name)
}
