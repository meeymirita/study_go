package main

import "fmt"

func main() {
	// Создание переменных через длинную запись
	number := 15
	text := "Hello"
	drob := 11.25
	boolean := true

	fmt.Println("number:", number)
	fmt.Println("text:", text)
	fmt.Println("drob:", drob)
	fmt.Println("boolean:", boolean)

	// Создание переменных через длинную запись
	var number1 int = 15
	var text1 string = "Hello"
	var drob1 float64 = 11.25
	var boolean1 bool = true

	fmt.Println("number1:", number1)
	fmt.Println("text1:", text1)
	fmt.Println("dro1:", drob1)
	fmt.Println("boolean1:", boolean1)

	var number2 int
	var text2 string
	var drob2 float64
	var boolean2 bool

	fmt.Println("number2:", number2)
	fmt.Println("text2:", text2)
	fmt.Println("drob2:", drob2)
	fmt.Println("boolean2:", boolean2)
}
