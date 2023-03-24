package main

import "fmt"

func main() {
		fmt.Println("go" + "lang") // String concatenation

		fmt.Println("1 + 1 =", 1+1) // Adding two integers returns an integer
		fmt.Println("7.0 / 3.0 =", 7.0/3.0) // Adding two floats returns a float

		// and, or, not operators are the same in Go as in javascript
		fmt.Println(true && false)
		fmt.Println(true || false)
		fmt.Println(!true)

		fmt.Println(len("hello")) // len() returns the length of a string
}