package main

import "fmt"

func main() {
	// инициализация при создании
	var user map[string]string = map[string]string{
		"name":     "Vasily",
		"lastName": "Romanov",
	}

	// сразу с нужной ёмкостью
	profile := make(map[string]string, 10)

	// количество элементов
	mapLength := len(user)

	fmt.Println("%d %+v\n", mapLength, profile)

	// если ключа нет - вернёт значение по умолчанию для типа
	mName := user["middleName"]
	fmt.Println("mName:", mName)

	// проверка существование ключа
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// пустая переменная ө только проверяем что ключ есть
	_, mNameexist2 := user["middleName"]
	fmt.Println("mNameExist2", mNameexist2)

	// удаление ключа
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}
