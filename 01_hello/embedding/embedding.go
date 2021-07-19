package main

import "fmt"

type Processor struct {
	processorName string
	cores         int
}

type Memory struct {
	memoryCapacity int
	momoryType     string
}

type Computer struct {
	Processor
	Memory
	price int
}

func main() {
	computer := Computer{}
	computer.price = 50000
	computer.processorName = "intel I5"
	computer.cores = 6
	computer.memoryCapacity = 8
	computer.momoryType = "DDr4"
	fmt.Println(computer)

}
