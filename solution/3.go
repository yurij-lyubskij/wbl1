package solution

import (
	"fmt"
	"runtime"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….)
//с использованием конкурентных вычислений.

func Sum(ch2 <-chan int, wg *sync.WaitGroup) {
	//после выполнения уменьшаем счетчик
	defer wg.Done()
	//начальное значение суммы
	sum := 0
	//итерируемся по каналу
	for num := range ch2 {
		//берем квадрат, добавляем к сумме
		sum += num
	}
	//выводим сумму
	fmt.Println(sum)
}

func Worker(ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
	//после выполнения уменьшаем счетчик
	defer wg.Done()
	//итерируемся по каналу
	for num := range ch1 {
		//берем число, считаем квадрат,
		//кладем квадрат в другой канал
		ch2 <- num * num
	}
}

func sumSquares(arr []int) {
	//для того, чтобы горутины конкурировали,
	//запустим все горутины на 1 ядре
	runtime.GOMAXPROCS(1)

	//используем WaitGroup, чтобы ждать завершения горутин
	wg := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	//создаем канал для входных
	//и выходных данных
	ch1 := make(chan int)
	ch2 := make(chan int)

	//задаем максимальное число горутин
	goroutinesNum := len(arr)
	if len(arr) > MaxRoutines {
		goroutinesNum = MaxRoutines
	}

	//запускаем горутины
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик
		wg.Add(1)
		go Worker(ch1, ch2, wg)
	}

	//увеличиваем 2 счетчик
	wg2.Add(1)
	//считаем сумму квадратов в горутине
	go Sum(ch2, wg2)

	//передаем входные данные в канал, из которого
	//читают горутины
	for _, j := range arr {
		ch1 <- j
	}
	//закрываем канал 1
	close(ch1)
	//ждем завершения воркеров
	wg.Wait()
	//закрываем канал 2, когда отработали горутины-воркеры
	close(ch2)
	wg2.Wait()
	//ждем завершения горутины-сумматора
}

func (*Index) N3() {
	//Отступ
	fmt.Println()
	//входной массив
	arr := [...]int{2, 4, 6, 8, 10}
	//запускаем нашу функцию
	sumSquares(arr[:])
}
