package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addcontinue(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(addcontinue(42, 13))
}
