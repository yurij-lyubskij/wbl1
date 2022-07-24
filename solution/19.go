package solution

import "fmt"

//Разработать программу, которая переворачивает
//подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

//Поворачиваем строку
func reverseString(in string) string {
	//Преобразуем в слайс байт
	input := []rune(in)
	//Выделяем память
	output := make([]rune, len(input))
	//В цикле пишем строку в новый слайс
	//в обратном порядке
	for i := 0; i < len(input); i++ {
		j := len(input) - i - 1
		output[i] = input[j]
	}
	//возвращаем строку из слайса
	return string(output)
}

func (*Index) N19() {
	//входные данные -строка
	fmt.Println("Разворачиваем строки")
	str := "главрыба - абырвалг"
	fmt.Println("Исходная строка :", str)
	//поворачиваем
	str = reverseString(str)
	//выводим результат
	fmt.Println("Повернутая строка :", str)
	fmt.Println("Так и должно быть)")
}
