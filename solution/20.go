package solution

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

//Поворачиваем строку
func reverseWords(in string, sep string) string {
	//Преобразуем в слайс слов
	input := strings.Split(in, sep)
	//Выделяем память
	output := make([]string, len(input))
	//В цикле пишем слова в новый слайс
	//в обратном порядке
	for i := 0; i < len(input); i++ {
		j := len(input) - i - 1
		output[i] = input[j]
	}
	//склеиваем строку из слайса
	return strings.Join(output, sep)
}

func (*Index) N20() {
	fmt.Println("Разворачиваем слова")
	//входные данные -строка
	str := "snow dog sun"
	fmt.Println("Исходная строка :", str)
	//поворачиваем
	str = reverseWords(str, " ")
	//выводим результат
	fmt.Println("Слова в обратном порядке :", str)
}
