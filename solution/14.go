package solution

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме
//способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func RTTI(p interface{}) {
	//просто кастим интерфейс
	switch p.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan interface {}")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	default:
		fmt.Println("Что-то новое!")
	}
}

func (*Index) N14() {
	//входные данные
	fmt.Println("входные данные:")
	A := 1
	B := "str"
	C := true
	D := make(chan interface{})
	fmt.Printf("A = %v, B = %v, C = %v, D = %v\n", A, B, C, D)
	fmt.Println("Определяем тип 2 разными способами")
	fmt.Println("1 способ -  свич по типу")
	//проверяем и выводим результат
	RTTI(A)
	RTTI(B)
	RTTI(C)
	RTTI(D)
	//также можно чере reflect.TypeOf()
	//тогда получаем значение reflect.Type,
	//c которым можно работать
	fmt.Println("2 способ - reflect.TypeOf")
	fmt.Println(reflect.TypeOf(A))
	fmt.Println(reflect.TypeOf(B))
	fmt.Println(reflect.TypeOf(C))
	fmt.Println(reflect.TypeOf(D))
}
