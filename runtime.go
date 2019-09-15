package main

import (
	"fmt"
	"runtime"
)

func test() {
	_, file, line, ok := runtime.Caller(1)
	fmt.Print(file)
	fmt.Print(line)
	fmt.Print(ok)
}

func main() {
	test()
}
