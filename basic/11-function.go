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

	fmt.Println("> closure")
	var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }
    var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers)
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	fmt.Println("> immediately-invoked function expression (IIFE)")
	var numbers_1 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
    var newNumbers = func(min int) []int {
        var r []int
        for _, e := range numbers_1 {
            if e < min {
                continue
            }
            r = append(r, e)
        }
        return r
    }(3)

    fmt.Println("original number :", numbers_1)
    fmt.Println("filtered number :", newNumbers)

    fmt.Println("> closure as return")
	var max_1 = 3
    var numbers_2 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
    var howMany, getNumbers = findMax(numbers_2, max_1)
    var theNumbers = getNumbers()

    fmt.Println("numbers\t:", numbers_2)
    fmt.Printf("find \t: %d\n\n", max)

    fmt.Println("found \t:", howMany)    // 9
    fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]

	fmt.Println("> function as params")
	var data = []string{"wick", "jason", "ethan"}
    var dataContainsO = filter(data, func(each string) bool {
        return strings.Contains(each, "o")
    })
    var dataLenght5 = filter(data, func(each string) bool {
        return len(each) == 5
    })
    fmt.Println("data asli \t\t:", data)
    fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
    fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)

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

func findMax(numbers []int, max int) (int, func() []int) {
    var res []int
    for _, e := range numbers {
        if e <= max {
            res = append(res, e)
        }
    }
    return len(res), func() []int {
        return res
    }
}

// alias cloure
type FilterCallback func(string) bool
func filter(data []string, callback FilterCallback) []string {
    var result []string
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}