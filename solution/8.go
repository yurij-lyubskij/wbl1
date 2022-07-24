package solution

import "fmt"

//Дана переменная int64. Разработать программу,
//которая устанавливает i-й бит в 1 или 0.

//функция для установки i бита в числе number
//в 1, если isOne - true, иначе - 0
func SetBit(number int64, i int, isOne bool) int64 {
	//0x000...1...000
	var mask int64 = 1 << i
	if !isOne {
		//и не
		number = number &^ mask
	}

	if isOne {
		//или
		number = number | mask
	}
	return number
}

func (*Index) N8() {
	//запускаем нашу функцию
	var number int64 = 0
	//отступ
	fmt.Println()
	//3 бит 0 в 1 - 8
	newNumber := SetBit(number, 3, true)
	fmt.Println(newNumber)
	//3 бит 8 в 0 - 0
	fmt.Println(SetBit(newNumber, 3, false))
}
