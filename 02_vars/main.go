package main

import "fmt"

// var name = "myname"

func main() {
	// var name = "myname"
	var age int = 37
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "myname"
	// size := 1.3
	// email := "myname@gmail.com"
	name, email := "myname", "myname@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
