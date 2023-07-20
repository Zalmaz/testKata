package main

import "fmt"

func main() {
	var a, b, result int
	var s string
	fmt.Scan(&a)
	fmt.Scan(&s)
	fmt.Scan(&b)
	switch s {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	fmt.Printf("%d %s %d = %d\n", a, s, b, result)
}
