package main

import "fmt"

func main() {
	fmt.Println("-> Array")

	// basic
	fmt.Println("> basic")
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	fmt.Println(names[0], names[1], names[2], names[3])

	// basic
	fmt.Println("> initial first value")
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Total element \t\t", len(fruits))
	fmt.Println("Fill all element \t", fruits)

	// multi dimensional array
	fmt.Println("> multi dimensional array")
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// looping array
	fmt.Println("> looping array")
	fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
	fmt.Println("> make array")
	var fruits_2 = make([]string, 2)
	fruits_2[0] = "apple"
	fruits_2[1] = "manggo"
	fmt.Println(fruits_2)  // [apple manggo]

	fmt.Println()
}