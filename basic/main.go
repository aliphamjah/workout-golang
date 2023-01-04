package main

// build in packpage
import "fmt"

func main() {
	helloWord()
	variable()
}

func helloWord()  {
	var firstName string = "alip"
	var lastName string
	lastName = "hamjah"
	greeting := "Selamat Pagi"

    	fmt.Println("-> Hello Word")
    	fmt.Println("Hello ", firstName, lastName, ",", greeting)
    	fmt.Println("\n")
}

func variable()  {
	// multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// underscore variable
	name, _ := "john", "wick"

	fmt.Println("-> variable")
	fmt.Println(first, second, third)
	fmt.Println(one, isFriday, twoPointTwo, say)
	fmt.Println(name)
	fmt.Println("\n")
}