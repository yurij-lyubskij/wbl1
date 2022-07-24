package solution

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

//спим указанное количество секунд
func sleep(seconds int) {
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	//блокируемся, пока таймер не сработает и в канале не появится сообщение
	_ = <-timer.C
}

func (*Index) N25() {
	//Отступ
	fmt.Println()
	//выводим что-то
	fmt.Println("Hello, ")
	//спим 5 с
	sleep(5)
	//выводим еще что-то
	fmt.Println("World")
}
