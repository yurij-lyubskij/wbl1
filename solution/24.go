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

type Point interface {
	Distance(anotherPoint Point) float64
}

type point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	var newPoint Point = &point{x: x, y: y}
	return newPoint
}

func (p *point) Distance(anotherPoint Point) float64 {
	point, ok := anotherPoint.(*point)
	if ok != true {
		//не тот тип, расстояние не посчитать
		return -1
	}
	square1 := math.Pow(float64(p.x-point.x), 2.0)
	square2 := math.Pow(float64(p.y-point.y), 2.0)
	return math.Sqrt(square1 + square2)
}

func N24() {
	//Отступ
	fmt.Println()
	//входные данные - координаты x1, y1, x2, y2
	x1, y1, x2, y2 := 1, 2, 4, 6
	fmt.Println(x1, y1, x2, y2)
	//создаем точки конструктором
	//X1, X2 имеют тип Point (а не point)
	X1 := NewPoint(x1, y1)
	X2 := NewPoint(x2, y2)
	//считаем расстояние, вызывая метод интерфейса
	//и выводим результат
	D := X1.Distance(X2)
	fmt.Println(D)
}
