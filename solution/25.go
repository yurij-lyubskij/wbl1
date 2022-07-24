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
	//выводим что-то
	fmt.Println("выводим что-то")
	//спим 5 с
	sleep(3)
	//выводим еще что-то
	fmt.Println("выводим еще что-то после того, как поспали")
}
