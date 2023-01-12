package main

import "fmt"

func main()  {
	fmt.Println("-> Looping")

	// basic
	fmt.Println("> basic")
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// for keyword
	fmt.Println("> for keyword")
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	// for keyword with break
	fmt.Println("> for keyword with break")
	i = 0
	for {
		fmt.Println("Angka", i)

		i++
		if i == 5 {
			break
		}
	}

	// break & continue keyword
	fmt.Println("> break & continue keyword")
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// loop on loop
	fmt.Println("> loop on loop")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// outerloop
	fmt.Println("> outerloop")
	outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 {
				break outerLoopLabel
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
	fmt.Println()
}