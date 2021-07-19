package main

import "fmt"

func main() {
	shoppingCart := map[string]int{
		"Keyboard": 100,
		"Mouse":    50,
		"laptop":   1000,
	}
	cart := shoppingCart["Mobile"]
	fmt.Println(shoppingCart["Mobile"])
	fmt.Println(cart)
	cart2, ok := shoppingCart["Mobile"]
	fmt.Println(shoppingCart["Mobile"])
	fmt.Println(cart2, ok)
	_, newok := shoppingCart["Mobile"] //you can use any variable in place of OK
	fmt.Println(shoppingCart["Mobile"])
	fmt.Println(newok)
	sc2 := shoppingCart
	fmt.Println(shoppingCart)
	fmt.Println(sc2)
	sc2["Mouse"] = 200
	fmt.Println(shoppingCart)
	fmt.Println(sc2)
	delete(shoppingCart, "laptop")
	fmt.Println(shoppingCart)
	delete(shoppingCart, "Mouse")
	fmt.Println(shoppingCart)
}
