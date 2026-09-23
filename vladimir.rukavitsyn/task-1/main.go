package main

import "fmt"

func main() {
	var (
		a, b     int
		operator string
	)
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("ERROR: Invalid first operand!!!")
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("ERROR: Invalid second operand!!!")
	}

	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("ERROR: Invalid operator!!!")
	}

	switch operator {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("ERROR: Division by zero!!!")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("ERROR: Invalid operator!!!")
	}

}
