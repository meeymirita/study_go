package main

import (
	"fmt"
)

func main() {
	score := 5
	if score > 10 {
		p("Ты красавчик")
	} else {
		p("Ты лошок")
	}

	//if err := ch.Qos(1, 0, false); err != nil {
	//	log.Fatal("Qos:", err)
	//}
}

func p(text string) {
	fmt.Println(text)
}
