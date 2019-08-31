package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// var i int 
	// i = 42
	i := 42
	// var i int = 42
	fmt.Println(i)

	var (
		a string = "a"
		// b string = "b"
	)
	fmt.Printf("a %T = \"%v\"", a, a)
}