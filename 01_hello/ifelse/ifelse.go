package main

import "fmt"

func main() {

	if i := 2; i == 3 {
		fmt.Println("this is simple if stateemnt")
	}
	shoppingCart := map[string]int{
		"Keyboard": 100,
		"Mouse":    50,
		"laptop":   1000,
	}
	fmt.Println(shoppingCart)
	if _, ok := shoppingCart["Mouse"]; ok {
		fmt.Println("Item is in Cart")
	}
}
