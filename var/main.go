package main

import (
	"fmt"
)

func main() {
	var a string = "Runoob"
	fmt.Println(a)
	var b bool = false
	fmt.Println(b)
	var c = 2.52
	fmt.Println(c)
	d := 50
	fmt.Println(d)
	e, f, g := true, false, 4.7
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	_, num, str := numbers()
	fmt.Println(num, str)
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
