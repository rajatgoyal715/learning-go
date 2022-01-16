package main

import "fmt"

// test function without any parameters
func test() {
	fmt.Println("test")
}

// test function with some parameters
func test(x int) {
	fmt.Println(x)
}

func main() {
	test() // this is how we call the function, can be re-used multiple times

	test(5)
}
