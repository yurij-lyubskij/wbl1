package main

import (
	s "l1/solution"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("Введите номер задачи")
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num < 1 || num > 26 {
		log.Fatalf("Нет такой задачи")
	}
	index := s.Index{}
	name := "N" + os.Args[1]
	reflect.ValueOf(&index).MethodByName(name).Call([]reflect.Value{})
}
