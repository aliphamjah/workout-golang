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
	arrayData()
	sliceData()
	mapData()
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

func arrayData() {
	fmt.Println("-> array")

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

func sliceData() {
	fmt.Println("-> slice")

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

func mapData() {
	fmt.Println("-> map")
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei",     chicken["mei"])     // mei 0

	fmt.Println("> initial value")
	// horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	// vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	fmt.Println(chicken1)
	fmt.Println(chicken2)
	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)

	fmt.Println("> iteration")
	var chicken6 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	for key, val := range chicken6 {
		fmt.Println(key, "  \t:", val)
	}

	fmt.Println("> delete")
	delete(chicken6, "januari")
	for key, val := range chicken6 {
		fmt.Println(key, "  \t:", val)
	}

	fmt.Println("> exist check")
	var value, isExist = chicken["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	fmt.Println("> slice & map")
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue",   "gender": "male"},
		map[string]string{"name": "chicken red",    "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	fmt.Println()
}