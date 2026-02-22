package main

import "fmt"

func main() {
	score := 12

	if score > 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Тебе нужно ещё многому научиться.")
	}

	score1 := 20

	if score1 > 15 {
		fmt.Println("Ты мега-красавчик!")
	} else if score1 > 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Тебе нужно ещё многому научиться.")
	}

	computerClub := true
	icecream := true

	// Я пойду гулять, если там либо будет компьютерный клуб
	// либо если там будет мороженное.
	// Если там не будет ни того ни другого, я никуда не пойду

	if computerClub || icecream {
		fmt.Println("Я иду гулять!")
	} else {
		fmt.Println("Я НЕ пойду гулять (")
	}

	sunny := true
	weekend := true

	// Я пойду гулять, только если будет солнечная погода и будет выходной,
	// иначе я никуда не пойду

	if sunny && weekend {
		fmt.Println("Я иду гулять!")
	} else {
		fmt.Println("Я НЕ пойду гулять (")
	}

	subscribed := true

	if !subscribed {
		fmt.Println("Я вижу ты не подписан! Подпишись, пожалуйста!")
	}
}
