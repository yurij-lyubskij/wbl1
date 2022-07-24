package solution

import "fmt"

//Имеется последовательность строк -
//(cat, cat, dog, cat, tree)
//создать для нее собственное множество.

// множество - просто мапа, где ключ - наш тип,
// а значение - пустая структура
func makeSet(Aslice []string) (A map[string]struct{}) {
	A = make(map[string]struct{}, 0)
	for _, key := range Aslice {
		A[key] = struct{}{}
	}
	return
}

func (*Index) N12() {
	//Отступ
	fmt.Println()
	//входные данные
	Aslice := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(Aslice)
	//ищем пересечение
	A := makeSet(Aslice)
	//выводим карту
	fmt.Println(A)
}
