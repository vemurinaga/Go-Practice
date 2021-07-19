package main

import "fmt"

func main() {
	shoppingCart := map[string]int{
		"Keyboard": 100,
		"Mouse":    50,
		"laptop":   1000,
	}
	fmt.Println(shoppingCart)
	shoppingCart["laptop"] = 2000
	fmt.Println(shoppingCart)
	fmt.Println(shoppingCart["Moblie"])
	cart2 := make(map[string]int)
	cart2 = map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(cart2)
}
