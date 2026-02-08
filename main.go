package main

import "fmt"

func main() {
	score := 0
	fmt.Println("Начало")
	fmt.Println("Ваш счёт:", score)

	plusOne(&score)
	fmt.Printf("Вы пролетели через %d трубу\n", score)
	fmt.Println("Ваш счёт:", score)
	plusOne(&score)
	fmt.Printf("Вы пролетели через %d трубу\n", score)
	fmt.Println("Ваш счёт:", score)
	plusOne(&score)
	fmt.Printf("Вы пролетели через %d трубу\n", score)
	fmt.Println("Ваш счёт:", score)

	fmt.Println("Вы разбились в тотал")
	fmt.Println("Фатал ерор")
}
func plusOne(score *int) {
	*score += 1
}
