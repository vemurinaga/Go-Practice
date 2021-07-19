package main

import "fmt"

type Student struct {
	name     string
	rollNo   int
	subjects []string
}

func main() {
	student1 := Student{
		name:   "chaitu",
		rollNo: 4,
		subjects: []string{
			"maths",
			"physics",
		},
	}
	fmt.Println(student1)
	student1.rollNo = 6
	fmt.Println(student1)
	student2 := Student{
		"aishini",
		04,
		[]string{
			"EVS",
			"VE",
		},
	}
	fmt.Println(student2)
}
