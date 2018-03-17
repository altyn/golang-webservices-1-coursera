package main

import "fmt"

// обычная функция
func doNothing() {
	fmt.Println("I'm regular function")
}

func main() {
	// анонимная функция
	func(in string) {
		fmt.Println("anon func out:", in)
	}("nobody")
	// return

	// присваивание анонимной функции в переменную
	printer := func(in string) {
		fmt.Println("pointer outs:", in)
	}
	printer("as variable")
	// return

	// определяем тип функции
	type strFuncType func(string)

	// функция коллбек
	worker := func(callback strFuncType) {
		callback("as callback")
	}
	worker(printer)
	// return

	// функция возвращает замыкание
	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}
	successLogger := prefixer("SUCCESS")
	successLogger("Expected behaviour")
}
