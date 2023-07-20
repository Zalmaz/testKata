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
}

func main() {
	var res int
	p := Operation{}
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
	p.operator = o[1]
	var err error
	p.firstNum, err = strconv.Atoi(o[0])
	if err != nil {
		panic(err)
	}
	p.lastNum, err = strconv.Atoi(o[2])
	if err != nil {
		panic(err)
	}

	switch p.operator {
	case "+":
		res = p.firstNum + p.lastNum
	case "-":
		res = p.firstNum - p.lastNum
	case "*":
		res = p.firstNum * p.lastNum
	case "/":
		res = p.firstNum / p.lastNum
	}
	fmt.Println(res)
}
