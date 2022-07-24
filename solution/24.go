package solution

import (
	"fmt"
	"math"
)

//
//Разработать программу нахождения
//расстояния между двумя точками,
//которые представлены в виде структуры Point
//с инкапсулированными параметрами x,y и конструктором.

//type Point interface {
//	Distance(anotherPoint Point) float64
//}

//type point struct {
//	x int
//	y int
//}

//x, y не видны за пределами пакета
type Point struct {
	x int
	y int
}

//конструктор
func NewPoint(x, y int) Point {
	var newPoint Point = Point{x: x, y: y}
	return newPoint
}

//метод нахождения расстояния
//считаем по теореме Пифагора
func (p *Point) Distance(point Point) float64 {
	square1 := math.Pow(float64(p.x-point.x), 2.0)
	square2 := math.Pow(float64(p.y-point.y), 2.0)
	return math.Sqrt(square1 + square2)
}

func (*Index) N24() {
	fmt.Println("входные данные - координаты точек 1 и 2")
	//входные данные - координаты x1, y1, x2, y2
	x1, y1, x2, y2 := 1, 2, 4, 6
	fmt.Printf("x1 = %d, y1 = %d, x2 = %d, y2 = %d\n", x1, y1, x2, y2)
	//создаем точки конструктором
	//X1, X2 имеют тип Point
	X1 := NewPoint(x1, y1)
	X2 := NewPoint(x2, y2)
	//считаем расстояние, вызывая метод
	//и выводим результат
	D := X1.Distance(X2)
	fmt.Println("Расстояние от 1 точки до 2 равно :", D)
}
