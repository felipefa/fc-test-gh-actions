package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
}

func sum(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}
