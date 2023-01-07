package main

// build in packpage
import "fmt"

func main() {
	helloWord()
	variable()
	typedata()
	constdata()
	operator()
	conditional()
	looping()
}

func helloWord()  {
	fmt.Println("-> Hello Word")

	var firstName string = "alip"
	var lastName string
	lastName = "hamjah"
	greeting := "Selamat Pagi"

    	fmt.Println("Hello ", firstName, lastName, ",", greeting)

	fmt.Println()
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

	fmt.Println()
}

func typedata()  {
	fmt.Println("-> type data")

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

func constdata()  {
	fmt.Println("-> const data")

	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	// type interface
	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	// multiple const
	const (
		square          = "kotak"
		isToday bool    = true
		sekarang
		numeric uint8   = 1
		floatNum        = 2.2
	)
	const three, four string = "tiga", "empat"
	fmt.Println()
}

func operator() {
	fmt.Println("-> operator")

	// ratio
	var value = (((2 + 6) % 3) * 4 - 2) / 3
	var isEqual = (value == 2)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// logic
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)

	fmt.Println()
}

func conditional()  {
	fmt.Println("-> conditional")
	var point = 8

	fmt.Println("> basic keyword")
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	fmt.Println("> temporary variable")

	var point_2 = 8840.0

	if percent := point_2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	fmt.Println("> switch case")

	var point_3 = 6

	switch point_3 {
		case 8:
			fmt.Println("perfect")
		case 7, 6, 5, 4:
			fmt.Println("awesome")
		default:
			{
				fmt.Println("not bad")
				fmt.Println("you can be better!")
			}
	}

	fmt.Println("> fallthrought")

	var point_4 = 6

	switch {
		case point_4 == 8:
			fmt.Println("perfect")
		case (point_4 < 8) && (point_4 > 3):
			fmt.Println("awesome")
			fallthrough
		case point_4 < 5:
			fmt.Println("you need to learn more")
		default:
			{
				fmt.Println("not bad")
				fmt.Println("you need to learn more")
			}
	}

	fmt.Println("> conditional nested")

	var point_5 = 10

	if point_5 > 7 {
		switch point_5 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point_5 == 5 {
			fmt.Println("not bad")
		} else if point_5 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point_5 == 0 {
				fmt.Println("try harder!")
			}
		}
	}

	fmt.Println()
}

func looping()  {
	fmt.Println("-> looping")

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