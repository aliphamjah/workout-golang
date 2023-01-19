package main

import "fmt"
import "strings"

type student struct {
    name  string
    grade int
}

func main()  {
	fmt.Println("> basic sample")
	var s1 = student{"john wick", 21}
    s1.sayHello()
    var name = s1.getNameAt(2)
    fmt.Println("nama panggilan :", name)

	fmt.Println()
	fmt.Println("> pointer method")
	var s2 = student{"john wick", 21}
    fmt.Println("s1 before", s2.name)
    // john wick
    s2.changeName1("jason bourne")
    fmt.Println("s1 after changeName1", s2.name)
    // john wick
    s2.changeName2("ethan hunt")
    fmt.Println("s1 after changeName2", s2.name)
    // ethan hunt
}

func (s student) sayHello() {
    fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName1(name string) {
    fmt.Println("---> on changeName1, name changed to", name)
    s.name = name
}

func (s *student) changeName2(name string) {
    fmt.Println("---> on changeName2, name changed to", name)
    s.name = name
}