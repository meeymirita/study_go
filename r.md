name := "Player" \
fmt.Printf("Игрок: %s\n", name)

fmt.Printf("Вы пролетели через %d трубу\n", score)
%d	int
%s	string
%f	float
%v	любой



- Функции

Если два или более последовательных именованных параметра функции имеют одинаковый тип, можно опустить тип для всех параметров, кроме последнего.
В этом примере мы сократили
x int, y int
до
x, y int

func add(x, y int) int {
return x + y
}

func main() {
fmt.Println(add(1, 2))
}

