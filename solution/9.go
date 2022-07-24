package solution

import (
	"fmt"
	"runtime"
	"sync"
)

//Разработать конвейер чисел. Даны два канала:
//в первый пишутся числа (x) из массива,
//во второй — результат операции x*2,
//после чего данные из второго канала
//должны выводиться в stdout.

func SquareOutput(ch2 <-chan int, wg *sync.WaitGroup) {
	//после выполнения уменьшаем счетчик
	defer wg.Done()
	//итерируемся по каналу
	for num2 := range ch2 {
		//выводим  num * 2
		fmt.Println(num2)
	}

}

func SquareCalc(ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
	//после выполнения уменьшаем счетчик
	defer wg.Done()
	//итерируемся по каналу
	for num := range ch1 {
		//берем число, считаем  num * 2
		//кладем  num * 2 в другой канал
		ch2 <- num * 2
	}
}

func SquareWrite(arr []int) {
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
		go SquareCalc(ch1, ch2, wg)
	}

	//запускаем горутины
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем 2 счетчик
		wg2.Add(1)
		//вывод квадратов в горутине
		go SquareOutput(ch2, wg2)
	}

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
	//ждем завершения горутин вывода
}

func N9() {
	//Отступ
	fmt.Println()
	//входной массив
	arr := [...]int{2, 4, 6, 8, 10, 12, 15, 18}
	//запускаем нашу функцию
	SquareWrite(arr[:])
}
