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
	//входные данные
	Aslice := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Исходный слайс строк", Aslice)
	//строим множество
	A := makeSet(Aslice)
	//выводим карту
	//можно было вывести только ключи
	fmt.Println("Получившееся множество", A)
}
