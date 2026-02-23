package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secret := rand.Intn(50)
	fmt.Println(secret)

	var guess int
	var count int

	for guess != secret {
		fmt.Print("Введи число: ")
		count++
		fmt.Scan(&guess)
		fmt.Println(check(secret, guess))
		fmt.Println("Попытки", count)
		if count == 7 {
			fmt.Println("Попытки кончились")
			break
		}
	}
}
func check(secret, guess int) string {

	if secret == guess {
		return "Ты угадал"
	}

	diff := secret - guess
	if diff <= 2 && diff >= -2 {
		return "близко"
	}
	if guess < secret {
		return "Число больше"
	}
	return "Число меньше"
}
