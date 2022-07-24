package solution

import (
	"errors"
	"fmt"
)

//Удалить i-ый элемент из слайса.

//проверяем, что есть такой номер элемента
//создаем новый слайс, добавляя к элементом до i
//все элементы после i
func DeleteFromSlice(slice []int, i int) ([]int, error) {
	if i < 0 || i >= len(slice) {
		return slice, errors.New("wrong element")
	}
	slice = append(slice[:i], slice[i+1:]...)
	return slice, nil
}

func (*Index) N23() {
	//Отступ
	fmt.Println()
	//входные данные - слайс
	slice := []int{7, 5, 6, 3, 8, 9}
	fmt.Println(slice)
	//удаляем элемент
	slice, err := DeleteFromSlice(slice, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	//выводим результат
	fmt.Println(slice)
}
