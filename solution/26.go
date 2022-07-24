package solution

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет,
//что все символы в строке уникальные
//(true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

func CheckUnique(input string) bool {
	RuneMap := make(map[rune]struct{})
	input = strings.ToLower(input)
	for _, oneRune := range input {
		_, ok := RuneMap[oneRune]
		if ok {
			return false
		}
		RuneMap[oneRune] = struct{}{}
	}
	return true
}

func N26() {
	//Отступ
	fmt.Println()
	//входные данные
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "aBcdb"

	fmt.Println(CheckUnique(s1))
	fmt.Println(CheckUnique(s2))
	fmt.Println(CheckUnique(s3))
	fmt.Println(CheckUnique(s4))
}
