package main

import (
	"fmt"
	"reflect"
)

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	a := [3]int{}
	b := make([]int, 3)
	fmt.Println("type of a =", reflect.TypeOf(a))
	fmt.Println("type of b =", reflect.TypeOf(b))

	fmt.Println(a, b)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("slice 1:", l)

	l = s[:5]
	fmt.Println("slice 2:", l)

	l = s[2:]
	fmt.Println("slice 3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i,_ := range twoD {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
				twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}