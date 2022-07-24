package solution

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

//интерфейс андроид
//можно заряжать проводом Type C
type android interface {
	insertTypeC()
}

//смартфон. Реализует
//интерфейс андроид
type Smartphone struct {
}

func (a *Smartphone) insertTypeC() {
	fmt.Println("Insert USB cable into type C")
}

//apple
//свой кабель - lightning
type apple struct{}

func (a *apple) insertLightning() {
	fmt.Println("Insert USB cable using Lightning Connector")
}

//адаптер
//содержит переходник на разъем lightning
type androidAdapter struct {
	connector *apple
}

//поэтому можно заряжать iPhone от TypeC
func (a *androidAdapter) insertTypeC() {
	a.connector.insertLightning()
}

type client struct {
}

//у клиента есть кабель TypeC
func (c *client) insertTypeC(usb android) {
	usb.insertTypeC()
}

func N21() {
	client := &client{}
	smartPhone := &Smartphone{}
	//клиент может заряжать смартфон без адаптера
	client.insertTypeC(smartPhone)
	appleConnector := &apple{}
	usbAdapter := &androidAdapter{
		connector: appleConnector,
	}
	//клиент может заряжать айфон, если есть адаптер
	client.insertTypeC(usbAdapter)
}
