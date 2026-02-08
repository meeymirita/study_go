package main

import "fmt"

func main() {
	text1 := "Get ready"
	score := 0
	text2 := "Game Over"
	fmt.Println(text1)
	fmt.Println("Ваш счёт:", score)
	fmt.Println("Вы пролете через первую трубу")
	plusOne(&score)
	fmt.Println("Ваш счёт:", score)
	fmt.Println("Вы пролете через вторую трубу")
	plusOne(&score)
	fmt.Println("Ваш счёт:", score)
	fmt.Println("Вы разбились в тотал")
	fmt.Println(text2)
}
func plusOne(score *int) {
	*score += 1
}
