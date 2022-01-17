package main

import "fmt"

func main() {
	toChange := "hello"
	var pointer *string = &toChange          // pointer stores the address of the string
	var pointerToPointer **string = &pointer // pointer to pointer stores the address of the pointer
	fmt.Println(pointer, pointerToPointer)
}
