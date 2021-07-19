package main

import "fmt"

const (
	i = iota
	j = iota
	k = iota
	a
)

const (
	b = iota + 1
	_
	c
	d
)

func main() {

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
