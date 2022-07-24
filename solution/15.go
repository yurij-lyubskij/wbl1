package solution

import (
	"fmt"
	"math/rand"
)

//К каким негативным последствиям может привести
//данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

// Во-первых, в памяти будет держаться вся большая строка,
//с которой мы не работаем. Во-вторых, может разрезаться руна
//в-третьих, скопируется не обязательно 100 символов

//генерация псевдослучайной строки
func createHugeString(size int64) string {
	A := []rune{}
	var i int64
	for i = 0; i < size; i++ {
		chars := []rune("ab⌘sdfdfменфв")
		randRune := chars[rand.Intn(len(chars))]
		A = append(A, randRune)
	}
	return string(A)
}

func copyString(input string, num int) (string, int) {
	//проверка, хватает ли длины исходной строки
	if num > len(input) {
		return input, len(input)
	}
	//выделяем память под временный слайс рун
	outSlice := make([]rune, num)
	//копируем из старого слайса в новый
	copy(outSlice, []rune(input)[:num])
	//возвращаем строку из слайса
	return string(outSlice), num
}

var justString string

func someFunc() {
	//создаем и выводи случайную строку
	v := createHugeString(1 << 10)
	fmt.Println(v)
	//временные переменные освободятся после выхода из функции
	//при использовании рун символы не обрезаются
	justString, _ = copyString(v, 100)
	fmt.Println(justString)
}

func N15() {
	someFunc()
}
