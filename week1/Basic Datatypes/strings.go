package main

func main() {
	// пустая строка по умолчанию
	var str string

	// со спец символами
	var hello string = "Привет\n\t"

	// без спец символов
	var world string = `Мир\n\t`

	// UTF-8 из коробки
	var helloworld = "Привет, Мир!"
	excercises := "Көнүгүүлөр"

	// одинарные кавычки для байт(uint8)
	var rawBinary byte = '\x27'

	// rune (uint32) для UTF-8 символов
	var someKyrgyz = "ҮҮҮӨӨ"

	helloworld := "Привет Мир"

	// конкатенация строк
	andGoodMorning := helloworld + " и Доброе утро!"

	// строки неизменяемы
	// can not assign to helloworld[0]
	helloworld[0] = 72

	// получение длины строки
	byteLen := len(helloworld)                  // 19 байт
	symbols := utf8.RuneCountString(helloworld) // 10 рун

	// полусение подстроки, в байтах, не в символах!
	hello := helloworld[:12] //Привет, 0-11 байты
	H := helloworld[0]       // byte, 72, не "П"

	// конвертация в слайс байт и обратно
	byteString = []byte(helloworld)
	helloWorld = string(byteString)
}
