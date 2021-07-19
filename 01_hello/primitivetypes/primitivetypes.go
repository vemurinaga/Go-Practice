package main

import (
	"fmt"
)

func main() {
	var foo bool = false
	i := 20
	j := 3
	fmt.Printf("%v, %T", foo, foo)
	fmt.Println(i + j)
	fmt.Println(i - j)
	fmt.Println(i / j)
	fmt.Println(i % j)
	fmt.Println(i * j)
	fmt.Println(i & j)
	fmt.Println(i | j)
	fmt.Println(i ^ j)
	fmt.Println(i &^ j)

}
