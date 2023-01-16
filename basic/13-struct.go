package main

import "fmt"

type student struct {
	name  string
	grade int
}

type person struct {
    name string
    age  int
}

type student_1 struct {
	grade int
	age  int
    person
}

type student_2 struct {
	grade int
    person
}

func main() {

	fmt.Println("> basic")
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)

	fmt.Println()
	fmt.Println("> initialize")
	var s1_1 = student{}
	s1_1.name = "wick"
	s1_1.grade = 2

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("student 1 :", s1_1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

	fmt.Println()
	fmt.Println("> pointer object")
	pointerObject()

	fmt.Println()
	fmt.Println("> embeded struct")
	embededStruct()

	fmt.Println()
	fmt.Println("> filling sub-struct")
	fillingSubStruct()

	fmt.Println()
	fmt.Println("> anonymous struct")
	anonymousStruct()

	fmt.Println()
	fmt.Println("> anonymous slice struct")
	anonymousSliceStruct()
}

func pointerObject() {
	var s1 = student{name: "wick", grade: 2}

	var s2 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)

	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
}

func embededStruct() {
	var s1 = student_1{}
    s1.name = "wick"
    s1.age = 21
    s1.person.age = 22
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("age   :", s1.age)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)
}

func fillingSubStruct() {
	var p1 = person{name: "john wick", age: 21}
	var s1 = student_2{person: p1, grade: 2}

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("grade :", s1.grade)
}

func anonymousStruct() {
	// anonymous struct without property
	var s1 = struct {
		person
        grade int
	}{}

	// anonymous struct with property
	var s2 = struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}

    s1.person = person{"wick", 21}
    s1.grade = 2
    s2.person = person{"john", 21}
    s2.grade = 3

    fmt.Println("1")
    fmt.Println("name  :", s1.person.name)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)
    fmt.Println("2")
    fmt.Println("name  :", s2.person.name)
    fmt.Println("age   :", s2.person.age)
    fmt.Println("grade :", s2.grade)
}

func anonymousSliceStruct()  {
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}
}