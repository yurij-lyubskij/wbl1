package solution

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет
//инкрементироваться в конкурентной среде. По завершению
//программа должна выводить итоговое значение счетчика.

//просто структура, но определены методы
type atomicCounter struct {
	atomicVal int32
}

//увеличивает на 1 атомарно
func (c *atomicCounter) Increment() {
	atomic.AddInt32(&c.atomicVal, 1)
}

//возвращает занчение атомарно
func (c *atomicCounter) ShowCounter() int32 {
	return atomic.LoadInt32(&c.atomicVal)
}

//конструктор
func NewCounter() *atomicCounter {
	return &atomicCounter{}
}

//воркер
func Counter(atomicVal *atomicCounter, wg *sync.WaitGroup) {
	//после выполнения уменьшаем счетчик wg
	defer wg.Done()
	//инкрементируем счетчик
	for i := 0; i < 10; i++ {
		atomicVal.Increment()
	}
}

func AtomicCount() {
	//для того, чтобы была параллельность
	//запустим на всех ядрах
	runtime.GOMAXPROCS(0)

	//используем WaitGroup, чтобы ждать завершения горутин
	wg := &sync.WaitGroup{}
	var atomicVal = NewCounter()
	//много горутин
	goroutinesNum := 120

	//запускаем горутины
	for i := 0; i < goroutinesNum; i++ {
		//увеличиваем счетчик wg
		wg.Add(1)
		go Counter(atomicVal, wg)
	}
	wg.Wait()
	//ShowCounter возвращает занчение атомарно
	fmt.Println("counter is = ", atomicVal.ShowCounter())
}

func (*Index) N18() {
	//Отступ
	fmt.Println()
	//запускаем нашу функцию
	AtomicCount()
}
