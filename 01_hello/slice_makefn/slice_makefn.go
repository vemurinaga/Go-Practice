package main

import "fmt"

func main() {
	var a []int = make([]int, 3)
	var b []int = make([]int, 4, 10)
	a = []int{1, 2, 3}
	var c []int = append(a, 5)
	var d []int = append(a[1:], 7) //append with slicing operator
	b = []int{5, 6, 7, 8}
	var e []int = append(b, a...)
	fmt.Printf("A: %v\n", a)
	fmt.Printf("B: %v\n", b)
	fmt.Printf("C: %v\n", c)
	fmt.Printf("D: %v\n", d)
	fmt.Printf("E: %v\n", e)
	fmt.Println(len(b))
	fmt.Println(cap(b))

}
