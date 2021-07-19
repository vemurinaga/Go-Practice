package main

import "fmt"

func main() {
	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
	fmt.Println(identityMatrix)
	fmt.Println(identityMatrix[0][0])
	identityMatrix[0][0] = 9
	fmt.Println(identityMatrix[0][0])

}
