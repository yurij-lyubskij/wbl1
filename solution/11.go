package solution

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.
func Intersection(A, B map[int]struct{}) map[int]struct{} {
	C := make(map[int]struct{})
	for key := range A {
		_, ok := B[key]
		if ok {
			C[key] = struct{}{}
		}
	}
	return C
}

func fillAB() (A, B map[int]struct{}) {
	A = make(map[int]struct{}, 0)
	B = make(map[int]struct{}, 0)
	Aslice := []int{1, 2, 3, 4, 5, 8, 9}
	Bslice := []int{2, 4, 6, 7, 9}
	for _, key := range Aslice {
		A[key] = struct{}{}
	}
	for _, key := range Bslice {
		B[key] = struct{}{}
	}
	return
}

func N11() {
	//Отступ
	fmt.Println()
	//заполняем множества
	A, B := fillAB()
	fmt.Println(A)
	fmt.Println(B)
	//ищем пересечение
	C := Intersection(A, B)
	//выводим карту
	fmt.Println(C)
}
