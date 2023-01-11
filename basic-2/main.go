package main

// build in packpage
import (
	"fmt"
	"strings"
	"math/rand"
	"math"
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

	fmt.Println("> return 2 value")
	var diameter float64 = 15
    var area, circumference = calculate(diameter)
	fmt.Printf("wide lingkaran\t\t: %.2f \n", area)
    fmt.Printf("circumference lingkaran\t: %.2f \n", circumference)

	fmt.Println("> predefined return value")
    area, circumference = calculate_1(diameter)
	fmt.Printf("wide lingkaran\t\t: %.2f \n", area)
    fmt.Printf("circumference lingkaran\t: %.2f \n", circumference)

	fmt.Println("> variadic")
	var avg = calculate_2(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
    var msg = fmt.Sprintf("Average : %.2f", avg)
    fmt.Println(msg)


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

func calculate(d float64) (float64, float64) {
    // count wide
    var area = math.Pi * math.Pow(d / 2, 2)
    // count circumference
    var circumference = math.Pi * d

    // return back 2 value
    return area, circumference
}

func calculate_1(d float64) (area float64, circumference float64) {
    area = math.Pi * math.Pow(d / 2, 2)
    circumference = math.Pi * d

    return
}

func calculate_2(numbers... int) float64 {
    var total int = 0
    for _, number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
    return avg
}