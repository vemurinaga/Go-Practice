package main

import "fmt"

func main() {
	var amounts [3]int = [3]int{10, 20, 30}
	amt := [3]int{30, 40, 50}
	a := amounts
	fmt.Printf("amount: %v\n", amounts)
	fmt.Printf("amount: %v\n", amt)
	fmt.Printf("%v\n", len(amt))
	a[0] = 51
	fmt.Printf("A: %v\n", a)
}
