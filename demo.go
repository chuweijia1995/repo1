package main

import "fmt"


// addNumbers computes the sum of two integers.
func addNumbers(a int, b int) int {
    return a + b
}

// main is the program entry point. It initializes two integers, computes their sum, and prints a descriptive message showing the operands and result.
func main() {
    num1 := 10
    num2 := 20

    sum := addNumbers(num1, num2)
    fmt.Println("The sum of", num1, "and", num2, "is:", sum)
}