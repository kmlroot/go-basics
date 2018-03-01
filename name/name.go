package name

import "fmt"

func GetName() string {
	var name string
	name = "Without name"

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)

	return name
}
