package main

import "fmt"

func main() {
	fmt.Print("Enter a Fahrenheit: ")
	var input float32

	fmt.Scanf("%f", &input)
	fmt.Println("Celsius:", (input-32)*(5.0/9.0))
}
