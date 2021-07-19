package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 12
	var j float32 = 3.4
	k := 3.4
	var foo string = string(i)
	var foo1 string = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", foo, foo)
	fmt.Printf("%v, %T\n", foo1, foo1)
	fmt.Println(i, j)
}
