package main

import "fmt"

func main() {
	var number1, number2, result float64
	var operator string
	fmt.Print("enter your number: ")
	fmt.Scanln(&number1)
	fmt.Print("enter another number: ")
	fmt.Scanln(&number2)
	fmt.Print("choose an operator (+, -, x, /): ")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "x":
		result = number1 * number2
	case "/":
		if number2 == 0 {
			println("Error: Division By Zero")
			return
		}
		result = number1 / number2
	default:
		fmt.Println("Invalid Operator")
		return
	}
	fmt.Printf("Result: %f\n", result)
}
