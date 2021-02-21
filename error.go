package main

import "fmt"

type Exception interface {
	getMessage() string
	getCode() int
	getLine() int
	getFile() string
}

type UserException struct {
}

func (userException UserException) getMessage() string {
	return "test"
}
func (userException UserException) getCode() int {
	return 200
}
func (userException UserException) getLine() int {
	return 12
}
func (userException UserException) getFile() string {
	return "dsf"
}

func main() {
	var exception Exception
	exception = new(UserException)
	defer func() {
		if e := recover(); nil != e {
			excep := e.(Exception)
			fmt.Println(excep.getMessage())
		}
	}()

	panic(exception)
}
