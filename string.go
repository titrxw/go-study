package main

import "fmt"

func main() {
	var test = "sdfsdgshgfhgf"
	fmt.Print(len(test))
	fmt.Print(string(test[0]))

	test = "括起来"
	fmt.Print(string(test[0]))
}
