package solution

import (
	"fmt"
	"runtime"
)

//Написать программу, которая конкурентно рассчитает
//значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

func Square(in chan int) {
	//итерируемся по каналу
	for num := range in {
		//считаем квадрат и выводим
		fmt.Print(num*num, " ")
	}

}

func concurrentSquares(arr []int) {
	//для того, чтобы горутины конкурировали,
	//запустим все горутины на 1 ядре
	runtime.GOMAXPROCS(1)

	//создаем канал
	ch := make(chan int)

	//задаем максимальное число горутин
	goroutinesNum := len(arr)
	if len(arr) > MaxRoutines {
		goroutinesNum = MaxRoutines
	}

	//запускаем горутины
	for i := 0; i < goroutinesNum; i++ {
		go Square(ch)
	}

	//передаем входные данные в канал, из которого
	//читают горутины
	for _, j := range arr {
		ch <- j
	}
}

func N2() {
	//входной массив
	arr := [...]int{2, 4, 6, 8, 10}
	//запускаем нашу функцию
	concurrentSquares(arr[:])
}
