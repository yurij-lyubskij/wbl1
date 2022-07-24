package solution

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func chanClose() {
	//для того, чтобы горутины конкурировали,
	//запустим все горутины на 1 ядре
	runtime.GOMAXPROCS(1)

	//используем WaitGroup, чтобы ждать завершения горутин
	wg := &sync.WaitGroup{}
	//создаем канал для входных данных
	ch1 := make(chan interface{})

	goroutinesNum := 10
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик
		wg.Add(1)
		// просто закроем канал
		go func() {
			//после выполнения уменьшаем счетчик
			defer wg.Done()
			//итерируемся по каналу
			for anything := range ch1 {
				fmt.Println(anything)
			}
		}()
	}
	//передаем случайные входные данные в канал, из которого
	//читают горутины
	//по таймеру просто закроем канал, и горутины выйдут из цикла и завершатся
	//ждем их завершения с wg.Wait()
	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(2 * time.Second)
	for _ = range ticker.C {
		select {
		//получаем сигнал
		case _ = <-timer.C:
			fmt.Println("graceful shutdown1")
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

func selectClose() {
	//для того, чтобы горутины конкурировали,
	//запустим все горутины на 1 ядре
	runtime.GOMAXPROCS(1)

	//используем WaitGroup, чтобы ждать завершения горутин
	wg := &sync.WaitGroup{}
	//создаем канал для входных данных
	ch1 := make(chan interface{})
	defer close(ch1)
	//создаем канал для сигнала о завершении
	ch2 := make(chan struct{})
	defer close(ch2)
	goroutinesNum := 10
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик
		wg.Add(1)
		go func() {
			//после выполнения уменьшаем счетчик
			defer wg.Done()
			for {
				select {
				//по сигналу выходим
				case _ = <-ch2:
					//выходим
					return
				case anything := <-ch1:
					fmt.Println(anything)
				}
			}
		}()
	}
	//передаем случайные входные данные в канал, из которого
	//читают горутины
	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(2 * time.Second)
	for _ = range ticker.C {
		select {
		//получаем сигнал
		case _ = <-timer.C:
			fmt.Println("graceful shutdown2")
			//посылаем пустую структуру столько раз, сколько горутин
			for i := 0; i < goroutinesNum; i++ {
				ch2 <- struct{}{}
			}
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

func ctxCancel() {
	//для того, чтобы горутины конкурировали,
	//запустим все горутины на 1 ядре
	runtime.GOMAXPROCS(1)

	//используем WaitGroup, чтобы ждать завершения горутин
	wg := &sync.WaitGroup{}
	//создаем канал для входных данных
	ch1 := make(chan interface{})
	defer close(ch1)
	//метод завершения горутин - с помощью контекста
	ctx, cancel := context.WithCancel(context.Background())
	goroutinesNum := 10
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик
		wg.Add(1)
		go func() {
			//после выполнения уменьшаем счетчик
			defer wg.Done()
			for {
				select {
				//идем сюда после вызова cancel
				case _ = <-ctx.Done():
					//выходим
					return
				case anything := <-ch1:
					fmt.Println(anything)
				}
			}
		}()
	}
	//передаем случайные входные данные в канал, из которого
	//читают горутины
	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(2 * time.Second)
	for _ = range ticker.C {
		select {
		//получаем сигнал
		case _ = <-timer.C:
			fmt.Println("graceful shutdown3")
			cancel()
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

func timeoutCancel() {
	//для того, чтобы горутины конкурировали,
	//запустим все горутины на 1 ядре
	runtime.GOMAXPROCS(1)

	//используем WaitGroup, чтобы ждать завершения горутин
	wg := &sync.WaitGroup{}
	//создаем канал для входных данных
	ch1 := make(chan interface{})
	defer close(ch1)
	//метод завершения горутин - с помощью таймаута на контексте
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	goroutinesNum := 10
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик
		wg.Add(1)
		go func() {
			//после выполнения уменьшаем счетчик
			defer wg.Done()
			for {
				select {
				//идем сюда после таймаута
				case _ = <-ctx.Done():
					//выходим
					return
				case anything := <-ch1:
					fmt.Println(anything)
				}
			}
		}()
	}
	//передаем случайные входные данные в канал, из которого
	//читают горутины
	ticker := time.NewTicker(time.Second)
	for _ = range ticker.C {
		select {
		//получаем сигнал
		case _ = <-ctx.Done():
			fmt.Println("graceful shutdown4")
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

func signalCancel() {
	//для того, чтобы горутины конкурировали,
	//запустим все горутины на 1 ядре
	runtime.GOMAXPROCS(1)

	//используем WaitGroup, чтобы ждать завершения горутин
	wg := &sync.WaitGroup{}
	//создаем канал для входных данных
	ch1 := make(chan interface{})
	defer close(ch1)
	//метод завершения горутин - с помощью сигнала
	//создаем канал для сигналов ОС
	goroutinesNum := 10
	sig := make(chan os.Signal, goroutinesNum)
	//слушаем сигналы SIGINT,SIGTERM
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик
		wg.Add(1)
		go func() {
			//после выполнения уменьшаем счетчик
			defer wg.Done()
			for {
				select {
				//идем сюда после сигнала
				case _ = <-sig:
					//выходим
					return
				case anything := <-ch1:
					fmt.Println(anything)
				}
			}
		}()
	}
	//передаем случайные входные данные в канал, из которого
	//читают горутины
	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(2 * time.Second)
	for _ = range ticker.C {
		select {
		//получаем сигнал
		case _ = <-timer.C:
			fmt.Println("graceful shutdown5")
			//посылаем сигнал на прекращение работы
			for i := 0; i < goroutinesNum; i++ {
				syscall.Kill(syscall.Getpid(), syscall.SIGINT)
				time.Sleep(100 * time.Millisecond)
			}
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

func (*Index) N6() {
	//запускаем наши функции
	chanClose()
	selectClose()
	ctxCancel()
	timeoutCancel()
	signalCancel()
}
