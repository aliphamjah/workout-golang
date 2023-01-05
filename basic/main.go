package main

// build in packpage
import "fmt"

func main() {
	helloWord()
	variable()
	typedata()
}

func helloWord()  {
	fmt.Println("-> Hello Word")

	var firstName string = "alip"
	var lastName string
	lastName = "hamjah"
	greeting := "Selamat Pagi"

    	fmt.Println("Hello ", firstName, lastName, ",", greeting)

	fmt.Println("\n")
}

func variable()  {
	fmt.Println("-> variable")

	// multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// underscore variable
	name, _ := "john", "wick"

	fmt.Println(first, second, third)
	fmt.Println(one, isFriday, twoPointTwo, say)
	fmt.Println(name)

	fmt.Println("\n")
}

func typedata()  {
	fmt.Println("-> variable")

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

	fmt.Println("\n")
}