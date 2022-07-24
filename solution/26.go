package solution

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет,
//что все символы в строке уникальные
//(true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

//Функция проверки
func CheckUnique(input string) bool {
	//Множество (Set) для всех всречавшихся рун
	RuneMap := make(map[rune]struct{})
	//приводим к нижнему регистру для
	//обеспечения регистронезависимости
	input = strings.ToLower(input)
	//итерируемся по строки, проверяя руны
	for _, oneRune := range input {
		//если руна уже встречалась, вернем false
		_, ok := RuneMap[oneRune]
		if ok {
			return false
		}
		//если новая, запишем в Set
		RuneMap[oneRune] = struct{}{}
	}
	//если дошли сюда, значит руны не повторились
	return true
}

func (*Index) N26() {
	//Отступ
	fmt.Println("Проверяем уникальность вхождения символов в строку")
	fmt.Println("Исходные строки")
	//входные данные - строки
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "aBcdb"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	//проверяем и выводим результат
	fmt.Println("проверяем и выводим результат по порядку")
	fmt.Println(CheckUnique(s1))
	fmt.Println(CheckUnique(s2))
	fmt.Println(CheckUnique(s3))
	fmt.Println(CheckUnique(s4))
}
