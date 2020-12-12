package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c
		d
	)
	fmt.Println(a, b, c, d)
}
