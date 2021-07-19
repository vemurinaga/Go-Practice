package main

import "fmt"

func main() {
	var amounts []int = []int{10, 20, 30}
	amt := []int{30, 40, 50}
	var b []int = amt
	fmt.Printf("amount: %v\n", amounts)
	fmt.Printf("amount: %v\n", amt)
	b[2] = 90
	fmt.Printf("amount: %v\n", amt)
	fmt.Printf("amount: %v\n", b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

}
