package solution

import (
	"fmt"
	"math/big"
)

//Разработать программу,
//которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20.

//умножаем 2 числа, используя пакет big
func Multiply(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Mul(a, b)
}

//делим 2 числа, используя пакет big
func Divide(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Quo(a, b)
}

//складываем 2 числа, используя пакет big
func Add(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Add(a, b)
}

//вычитаем одно число из другого, используя пакет big
func Subtract(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Sub(a, b)
}

func N22() {
	//задаем большие числа строкой
	aString := "123456789101112131415161700000000.0003"
	bString := "567891011121314151617000000000000.0007"
	//создаем переменные, используя фабричные функции
	a, ok := new(big.Float).SetString(aString)
	if ok != true {
		fmt.Println("Not a number")
	}
	b, ok := new(big.Float).SetString(bString)
	if ok != true {
		fmt.Println("Not a number")
	}
	//Выводим значения и результат арифметических операций
	fmt.Printf("a = %v, b = %v\n", a, b)
	fmt.Printf("a * b = %v\n", Multiply(a, b))
	fmt.Printf("a / b = %v\n", Divide(a, b))
	fmt.Printf("a + b = %v\n", Add(a, b))
	fmt.Printf("a - b = %v\n", Subtract(a, b))
}
