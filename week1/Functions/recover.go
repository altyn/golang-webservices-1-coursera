package main

import (
	"fmt"
)

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend FIRST:", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend SECOND:", err)
			// panic(2)
		}
	}()
	fmt.Println("Some useful work")
	panic("Something bad happend")
	return
}

func main() {
	deferTest()
	return
}
