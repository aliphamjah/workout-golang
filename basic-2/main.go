package main

// build in packpage
import (
	"fmt"
	"strings"
	"math/rand"
)

func main() {
	function()
}

func function() {
	fmt.Println("-> Function")

	fmt.Println("> basic")
	var names = []string{"John", "Wick"}
	printMessage("halo", names)

	fmt.Println("> return")
	var randomValue int

    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)

	fmt.Println()
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min + 1) + min
	return value
}