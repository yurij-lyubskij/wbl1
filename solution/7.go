package solution

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

//Структура мапы
type ConcurMap struct {
	mu   *sync.RWMutex
	cMap map[int]int
}

//Сохранение значение с использованием мютекса
func (c *ConcurMap) Store(key int, value int) {
	c.mu.Lock()
	c.cMap[key] = value
	c.mu.Unlock()
}

//Забираем занчение, мютекс на чтение
func (c *ConcurMap) Load(key int) (val int, ok bool) {
	c.mu.RLock()
	val, ok = c.cMap[key]
	c.mu.RUnlock()
	return val, ok
}

//Создание ссылки на новую мапу
func NewMap() *ConcurMap {
	return &ConcurMap{
		mu:   &sync.RWMutex{},
		cMap: make(map[int]int),
	}
}

//пишем либо читаем в зависимости от четности
func StoreLoad(myMap *ConcurMap, ch1 <-chan int, wg *sync.WaitGroup) {
	//после выполнения уменьшаем счетчик
	defer wg.Done()
	//итерируемся по каналу
	for i := range ch1 {
		if i%2 == 0 {
			myMap.Store(i/2, i)
		}
		if i%2 == 1 {
			val, ok := myMap.Load(i / 2)
			if ok != true {
				fmt.Println("no such key found")
			}
			fmt.Printf("key = %d, val = %d\n", i/2, val)
		}
	}
}

//проверка - чтение и запись в горутинах
func MapCheck() {
	myMap := NewMap()
	//используем WaitGroup, чтобы ждать завершения горутин
	wg := &sync.WaitGroup{}
	//создаем канал для входных данных
	ch1 := make(chan int)

	goroutinesNum := 10
	//запускаем горутины
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик
		wg.Add(1)
		//не передаем контекст, т.к. просто закроем канал по таймауту
		go StoreLoad(myMap, ch1, wg)
	}

	//исходные данные
	for i := 0; i < 2*goroutinesNum; i++ {
		ch1 <- i
	}

	//закрываем канал 1
	close(ch1)
	//ждем завершения воркеров
	wg.Wait()
}

func N7() {
	//запускаем нашу функцию
	MapCheck()
}
