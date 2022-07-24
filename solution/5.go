package solution

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно
//отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

//воркер-читатель
func ReaderTimed(ch1 <-chan interface{}, wg *sync.WaitGroup) {
	//после выполнения уменьшаем счетчик
	defer wg.Done()
	//итерируемся по каналу
	for anything := range ch1 {
		fmt.Println(anything)
	}
}

func readingWithTimer() {
	//для того, чтобы горутины конкурировали,
	//запустим все горутины на 1 ядре
	runtime.GOMAXPROCS(1)

	//используем WaitGroup, чтобы ждать завершения горутин
	wg := &sync.WaitGroup{}
	//создаем канал для входных данных
	ch1 := make(chan interface{})

	//считываем число горутин из конфига
	err := ReadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	goroutinesNum := viper.GetInt("goroutines")

	//считываем число cекунд до остановки из конфига
	seconds := viper.GetInt("seconds")
	//получаем время задержки
	waitTime := time.Duration(seconds) * time.Second
	//получаем контекст с таймаутом
	ctx, _ := context.WithTimeout(context.Background(), waitTime)

	//запускаем горутины
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик
		wg.Add(1)
		//не передаем контекст, т.к. просто закроем канал по таймауту
		go ReaderTimed(ch1, wg)
	}

	//передаем случайные входные данные в канал, из которого
	//читают горутины по таймеру
	ticker := time.NewTicker(time.Second)
	for _ = range ticker.C {
		select {
		//получаем таймаут
		//	можно было сделать по таймеру
		//timer := time.NewTimer(2*time.Second)
		//t := <-timer.C
		//или передавать контекст, и сделать
		//select в горутине
		//но канал должен закрывать producer
		case <-ctx.Done():
			fmt.Println("graceful shutdown after timeout")
			//закрываем канал 1
			close(ch1)
			//ждем завершения воркеров
			wg.Wait()
			//выходим
			return
		default:
			//просто пишем в канал
			j := rand.Intn(100)
			ch1 <- j
		}
	}
}

func N5() {
	//запускаем нашу функцию
	readingWithTimer()
}
