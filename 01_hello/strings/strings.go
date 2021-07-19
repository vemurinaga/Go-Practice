package main

import (
	"fmt"
)

func main() {
	var string1 string = "this is my first string"
	string2 := 'T'
	s1 := []byte(string1)
	//s2 := []byte(string2)
	fmt.Printf("%v, %T\n", string1, string1)
	fmt.Printf("%v\n", string1[1])
	fmt.Printf("%v, %T/n", string(string1[1]), string1)
	fmt.Printf("v, %T/n", s1, s1)
	fmt.Printf("%v, %T", string2, string2)
}
