package main

import "fmt"

const (
	User    string = "Admin"
	Producr string = "mobile"
)

const incr = iota

func main() {
	const i int = 12
	const j float32 = 3.14

	fmt.Println(i, j)
}
