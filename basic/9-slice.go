package main

import "fmt"

func main() {
	fmt.Println("-> Slice")

	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits) // ["apple", "grape"]

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// update value
	fmt.Println("> update value")
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// lev and cap
	fmt.Println("> lev and cap")
	fmt.Println(len(fruits))  // len: 4
	fmt.Println(cap(fruits))  // cap: 4

	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	// append
	fmt.Println("> append")
	var cFruits = append(fruits, "papaya")
	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	// copy
	fmt.Println("> copy")
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)
	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	// slice with cap
	fmt.Println("> slice with cap")
	var fruits_2 = []string{"apple", "grape", "banana"}
	var aFruits_fruits_2 = fruits_2[0:2]
	var bFruits_fruits_2 = fruits_2[0:2:2]

	fmt.Println(fruits_2)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits_2)) // len: 3
	fmt.Println(cap(fruits_2)) // cap: 3

	fmt.Println(aFruits_fruits_2)      // ["apple", "grape"]
	fmt.Println(len(aFruits_fruits_2)) // len: 2
	fmt.Println(cap(aFruits_fruits_2)) // cap: 3

	fmt.Println(bFruits_fruits_2)      // ["apple", "grape"]
	fmt.Println(len(bFruits_fruits_2)) // len: 2
	fmt.Println(cap(bFruits_fruits_2)) // cap: 2


	fmt.Println()
}