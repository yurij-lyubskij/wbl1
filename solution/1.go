package solution

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action
//от родительской структуры Human (аналог наследования).

import "fmt"

//родительская структура
type Human struct {
	name string
	age  int
}

//метод родительской структуры
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s \n", h.name)
}

//структура, встраивающая Human
type Action struct {
	Human
	profession string
}

func (*Index) N1() {
	//создаем структуру, содержащую втроенную структуру Human
	fmt.Println("создаем структуру, содержащую втроенную структуру Human")
	s := Action{Human{name: "Ivan", age: 18}, "student"}
	fmt.Println(s)
	//Вызываем метод встроенной структуры 2 способами
	fmt.Println("Вызываем метод встроенной структуры 2 способами")
	s.SayHi()
	s.Human.SayHi()
}
