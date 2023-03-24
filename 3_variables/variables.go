package main // Every Go program must start with a package declaration

import (
	"fmt" // Importing the fmt package
	"reflect" // Importing the reflect package
)

func main() { // Every Go program must have a main() function
	var a = "initial" // Declaring a variable with the var keyword
	fmt.Println(a) // Printing the variable

	var b, c int = 1, 2 // Declaring multiple variables
	fmt.Println(b, c)

	var d = true // Declaring a variable without specifying the type
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d)) // Printing the type of the variable

	var e int // Declaring a variable without initializing it
	fmt.Println(e)
	
	f := "apple" // Declaring a variable using the short variable declaration operator
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))
}