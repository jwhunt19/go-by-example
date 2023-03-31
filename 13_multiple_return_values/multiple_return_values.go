package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func test(x, y, z int) (int, int) {
	return (x * y) + z, (z * y) + x
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	d, e := test(2, 6, 7)
	fmt.Println(d)
	fmt.Println(e)
}