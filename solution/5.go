package solution

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные
//данные из канала и выводят в stdout. Необходима возможность
//выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров

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

//воркер-читатель
func ReaderTimed(ctx context.Context, ch1 <-chan interface{}, wg *sync.WaitGroup) {
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
	seconds := viper.GetInt("seconds")
	waitTime := time.Duration(seconds) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), waitTime)

	//запускаем горутины
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик
		wg.Add(1)
		go ReaderTimed(ctx, ch1, wg)
	}

	//передаем случайные входные данные в канал, из которого
	//читают горутины по таймеру
	ticker := time.NewTicker(time.Second)
	for _ = range ticker.C {
		select {
		//получаем сигнал
		case <-ctx.Done():
			fmt.Println("graceful shutdown after timeout")
			//закрываем канал 1
			close(ch1)
			//ждем завершения воркеров
			wg.Wait()
			//выходим
			return
		default:
			j := rand.Intn(100)
			ch1 <- j
		}
	}
}

func N5() {
	//запускаем нашу функцию
	readingWithTimer()
}
