package main

import(
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Print("Please enter a floating point number: ")
	
	fmt.Scanf("%f", &input)
	
	fmt.Print("The truncation of your number is: ", math.Trunc(input), "\n")
}