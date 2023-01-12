package main

import "fmt"

func main()  {
	fmt.Println("-> Type data")

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	fmt.Printf("positive number: %d\n", positiveNumber)
	fmt.Printf("negative number: %d\n", negativeNumber)

	var decimalNumber = 2.62
	fmt.Printf("desimave number: %f\n", decimalNumber)
	fmt.Printf("desimave number: %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist ? %t \n", exist)

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	fmt.Println()
}