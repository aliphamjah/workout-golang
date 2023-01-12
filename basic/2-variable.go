package main

import "fmt"

func main()  {
	fmt.Println("-> Variable")

	// multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// underscore variable
	name, _ := "john", "wick"

	fmt.Println(first, second, third)
	fmt.Println(one, isFriday, twoPointTwo, say)
	fmt.Println(name)

	fmt.Println()
}