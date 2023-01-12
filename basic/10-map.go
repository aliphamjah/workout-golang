package main

import "fmt"

func main() {
	fmt.Println("-> Map")
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