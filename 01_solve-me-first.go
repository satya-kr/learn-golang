package main

import "fmt"

func solveMeFirst(a, b uint32) uint32 {
	return a + b
}

func main() {
	res  := solveMeFirst(1, 2)
	fmt.Println(res)
}
