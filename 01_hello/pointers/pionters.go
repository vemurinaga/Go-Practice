package main

import "fmt"

type Foo struct {
	bar int
}

func main() {
	var a int = 12
	var b *int = &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) //derefences
	var foo *Foo
	fmt.Println(foo)
	foo = new(Foo)
	fmt.Println(foo)
	fmt.Println(*foo)
	fmt.Println((*foo).bar)
	(*foo).bar = 45
	fmt.Println((*foo).bar)
	foo.bar = 54
	fmt.Println(foo.bar)

}
