package name

import (
	"fmt"
	"strings"
)

func GetName() string {
	var name string
	name = "Without name"

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)

	return name
}

func String2() {
	var text = "Mauricio Serna, Alba Florez, Hello Golang"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Alba", "Luz Albita", -1))
	fmt.Println(strings.Split(text, ","))
}
