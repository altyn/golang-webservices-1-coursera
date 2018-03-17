package main

import (
	"fmt"

	"./coursera/visibility/person"
)

func main() {
	p := person.NewPerson(1, "rvasily", "secret")

	// p.secret undefined (cannot refer to unexpected field or method secret)
	// fmt.Printf("main.PrintPerson: %+v\n", p.secret)

	secret := person.GetSecret(p)
	fmt.Println("Get Secret", secret)
}
