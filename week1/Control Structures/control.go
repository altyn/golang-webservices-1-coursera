package main

import "fmt"

func main() {
	// простое условие
	boolVal := true
	if boolVal {
		fmt.Println("BoolVal is true")
	}

	mapVal := map[string]string{"name": "rvasily"}
	// условие блока инициализации
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name = ", keyValue)
	}
	// получаем только признак существования ключа
	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println(" key 'name' exists")
	}

	cond := 1
	// множественное if else
	if cond == 1 {
		fmt.Println("cond is 1")
	} else if cond == 2 {
		fmt.Println("cond is 2")
	}

	// switch по 1 переменной
	strVal := "name"
	switch strVal {
	case "name":
		fallthrough
	case "test", "lastName":
		// some work
	default:
		// some work
	}

	// switch как замена многим ifelse
	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11:
		fmt.Println("first block")
	case val2 > 10:
		fmt.Println("second block")
	}

	// выход из цикла, находясь внутри цикла
Loop: // метка
	for key, val := range mapVal {
		println("switch in loop", key, val)
		switch {
		case key == "lastName":
			break
			println("don't pront this")
		case key == "firstName" && val == "Vasily":
			println("switch -break loop here")
			break Loop
		}
	} // конец for

}
