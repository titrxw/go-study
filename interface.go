package main

import "fmt"

type Phone interface {
	call()
}

type APhone struct {
}

func (aPhone APhone) call() {
	fmt.Print("we")
}

func main() {
	var mphone Phone
	mphone = new(APhone)
	mphone.call()
}
