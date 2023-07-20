package main

import "fmt"

func main() {
	var a, b, result int
	var s string
	fmt.Scan(&a)

	fmt.Scan(&s)
	fmt.Scan(&b)
	if a < 0 || b < 0 || a > 10 || b > 10 {
		panic("Ошибка, цифры не валидны")
	}
	switch s {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("Операнд не валиден")
	}
	fmt.Printf("%d %s %d = %d\n", a, s, b, result)
}
