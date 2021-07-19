package main

import "fmt"

type Student struct {
	name     string
	rollno   int
	subjects []string
}

func main() {
	student1 := Student{
		name:   "aishini",
		rollno: 04,
		subjects: []string{
			"evs",
			"ve",
		},
	}
	fmt.Println(student1)
	student2 := student1
	student2.name = "krithik"
	fmt.Println(student1)
	fmt.Println(student2)
	student3 := &student1
	student3.name = "chiaitanya"
	fmt.Println(student1)
	fmt.Println(student3)
}
