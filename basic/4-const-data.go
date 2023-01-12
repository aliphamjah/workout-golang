package main

import "fmt"

func main()  {
	fmt.Println("-> Const data")

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