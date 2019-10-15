package go_study

import "fmt"

func main() {
	var ass interface{} = 2
	value, ok := ass.(int)
	fmt.Print(value)
	fmt.Print(ok)

	as1 := 2
	va, err := interface{}(as1).(int)
	fmt.Print(va)
	fmt.Print(err)
}
