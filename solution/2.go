package solution

import (
	"fmt"
	"runtime"
)

//Написать программу, которая конкурентно рассчитает
//значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

const MaxRoutines = 10

func Square(in chan int) {
	num := <-in
	fmt.Print(num*num, " ")
}

func concurrentSquares(arr []int) {
	//для того, чтобы горутины конкурировали,
	//запустим все горутины на 1 ядре
	runtime.GOMAXPROCS(1)

	//создаем канал
	ch := make(chan int)
	goroutinesNum := len(arr)
	if len(arr) > MaxRoutines {
		goroutinesNum = MaxRoutines
	}
	for i := 0; i < goroutinesNum; i++ {
		go Square(ch)
	}

	for _, j := range arr {
		ch <- j
	}
}
func N2() {
	arr := [...]int{2, 4, 6, 8, 10}
	concurrentSquares(arr[:])
}
