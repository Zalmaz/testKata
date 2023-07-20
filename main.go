package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultOperator = "+-/*"
)

type Operation struct {
	firstNum int
	lastNum  int
	operator string
	isARab   bool
	result   int
}

func main() {
	p := Operation{}
	str := input()
	slice := split(str)
	validator(&p, slice)
	operation(&p)

	fmt.Println(p.result)
}

func input() string {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	str := s.Text()
	return str
}

func split(str string) []string {
	slice := strings.Split(str, " ")
	if len(slice) != 3 {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	return slice
}

func validator(o *Operation, slice []string) {
	var err error
	if !strings.Contains(defaultOperator, slice[1]) && len(slice) != 1 {
		panic("Невалидный оператор")
	}
	o.operator = slice[1]
	o.firstNum, err = strconv.Atoi(slice[0])
	if err != nil || o.firstNum > 10 || o.firstNum < 0 {
		panic("Неверное первое число")
	}
	o.lastNum, err = strconv.Atoi(slice[2])
	if err != nil || o.lastNum > 10 || o.lastNum < 0 {
		panic("Неверное второе число")
	}

}

func operation(o *Operation) {
	switch o.operator {
	case "+":
		o.result = o.firstNum + o.lastNum
	case "-":
		o.result = o.firstNum - o.lastNum
	case "*":
		o.result = o.firstNum * o.lastNum
	case "/":
		o.result = o.firstNum / o.lastNum
	}
}
