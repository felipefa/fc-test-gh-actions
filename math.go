package main

import "fmt"

func main() {
	fmt.Println(Sum(10, 10))
}

func Sum(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}
