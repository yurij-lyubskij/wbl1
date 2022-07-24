package solution

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.

//Определение значения элемента в середине структуры данных.
//Полученное значение сравнивается с ключом.
//Если ключ меньше значения середины,
//то поиск осуществляется в первой половине элементов, иначе — во второй.
//Поиск сводится к тому, что вновь определяется
//значение серединного элемента в выбранной половине и сравнивается с ключом.
//Процесс продолжается до тех пор, пока не будет найден
//элемент со значением ключа или не станет пустым интервал для поиска.

func BinarySearch(arr []int, key int) int {
	// определяются границы поиска
	low := 0
	high := len(arr) - 1

	for low <= high {
		// mid - значение в середине
		//между границами
		mid := (low + high) / 2
		midVal := arr[mid]

		if midVal < key {
			//сдвигаем левую границу
			low = mid + 1
		} else if midVal > key {
			//сдвигаем правую границу
			high = mid - 1
		} else {
			//midVal == key
			//нашли
			return mid
		}
	}
	// значение не найдено в массиве
	return -1
}

func (*Index) N17() {
	//Отступ
	fmt.Println()
	//входные данные -отсортированный массив
	arr := []int{-1, 1, 2, 3, 10, 15, 100, 500, 1000}
	//ищем
	a := BinarySearch(arr, -1)
	//выводим результат
	fmt.Println(a)
	//ищем
	a = BinarySearch(arr, 15)
	//выводим результат
	fmt.Println(a)
	//ищем
	a = BinarySearch(arr, 100)
	//выводим результат
	fmt.Println(a)
}
