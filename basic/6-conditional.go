package main

import "fmt"

func main()  {
	fmt.Println("-> Conditional")
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