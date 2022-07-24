package solution

import "fmt"

//Поменять местами два числа без создания временной переменной.

func (*Index) N13() {
	//входные данные
	fmt.Println("входные данные:")
	A := 1
	B := 2
	fmt.Printf("A = %d, B = %d\n", A, B)
	fmt.Println("просто меняем местами")
	//просто меняем местами
	A, B = B, A
	//выводим результат
	fmt.Printf("A = %d, B = %d\n", A, B)
}
