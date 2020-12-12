package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokia NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am Iphone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()
}
