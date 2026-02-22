package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello, World!")
	}

	// 🔋
	// 🐤

	score := 0

	fmt.Println("Get Ready")
	fmt.Println("Счёт:", score)
	fmt.Println("")

	for i := 1; i <= 10; i++ {
		fmt.Println("---------------------")
		fmt.Println("Вы подлетаете к трубе!", i)
		fmt.Println("🐤 🔋🔋")
		fmt.Println("")

		fmt.Println("Вы пролетаете через трубу!", i)
		fmt.Println("🔋🐤🔋")
		fmt.Println("")

		fmt.Println("Вы пролетели через трубу!", i)
		fmt.Println("🔋🔋 🐤")
		fmt.Println("")

		score++

		fmt.Println("Счёт:", score)
		fmt.Println("")
	}

	// 🔋
	// 🛢
	// 🐤

	fmt.Println("Начинаю генерировать уровень!")
	fmt.Println("")

	for i := 1; i <= 5; i++ {
		fmt.Println("Труба номер", i)
		fmt.Println("----")

		if i%2 == 0 {
			fmt.Println("🛢🛢")
		} else {
			fmt.Println("🔋🔋")
		}

		fmt.Println("----")
		fmt.Println("")
	}

	fmt.Println("Генерация уровня окончена!")

	// 🔋
	// 🐤
	// ❌

	score := 0

	fmt.Println("Get Ready")
	fmt.Println("Счёт:", score)
	fmt.Println("")

	// Бесконечный цикл (аналог while true в других языках)
	for {
		fmt.Println("---------------------")
		fmt.Println("Я подлетаю к трубе!")
		fmt.Println("🐤 🔋🔋")
		fmt.Println("")

		fmt.Println("Я пролетаю через трубу!")
		fmt.Println("🔋🐤🔋")
		fmt.Println("")

		if rand.Intn(8) == 1 {
			fmt.Println("Я врезался в трубу :(")
			fmt.Println("🔋❌🔋")

			// прерывание цикла
			break
		}

		fmt.Println("Я пролетел через трубу!")
		fmt.Println("🔋🔋 🐤")
		fmt.Println("")

		score++

		fmt.Println("Счёт:", score)
		fmt.Println("")

		time.Sleep(500 * time.Millisecond)
	}
}
