package main

import "fmt"

func main() {
	msg := "Hello user"
	writemessages(msg)
	fmt.Println(msg)
}

func writemessages(msg string) {
	fmt.Println(msg)
	msg = "Hello youtube"
}
