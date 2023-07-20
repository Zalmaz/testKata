package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	defaultOperator = "+-/*"
)

type Operation struct {
	firstNum int
	lastNum  int
	operator string
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	str := s.Text()
	o := strings.Split(str, " ")
	if len(o) != 3 {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	if !strings.Contains(defaultOperator, o[1]) && len(o) != 1 {
		panic("Невалидный оператор")
	}

}
