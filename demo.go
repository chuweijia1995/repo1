package main

import "fmt"


func addNumbers(a int, b int) int {
    return a + b
}

func main() {
    num1 := 10
    num2 := 20

    sum := addNumbers(num1, num2)
    fmt.Println("The sum of", num1, "and", num2, "is:", sum)
}
