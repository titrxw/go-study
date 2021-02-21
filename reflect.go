package main

import (
	"fmt"
	"reflect"
)

type UserEE struct {
}

//https://www.jianshu.com/p/db16d3bd5908
//在反射中要注意大小写,需要首字母大写（待查询）
func (this UserEE) GetMessage() string {
	return "test"
}
func (this UserEE) getCode() int {
	return 200
}
func (this UserEE) getLine() int {
	return 12
}
func (this UserEE) getFile() string {
	return "dsf"
}

type Blog struct {
}

func (this Blog) Test() string {
	fmt.Println("this is Test method")
	return "we"
}

func main() {
	var o = new(Blog)
	v := reflect.ValueOf(o)
	fmt.Println(v)
	m := v.MethodByName("Test")
	rets := m.Call(nil)
	fmt.Println(rets)
	fmt.Println(rets[0])

	var u = new(UserEE)
	p := reflect.ValueOf(u)

	m1 := p.MethodByName("GetMessage")
	fmt.Print(m1)
	res := m1.Call(nil)
	//args := []reflect.Value{reflect.ValueOf("fdsf"), reflect.ValueOf(12)}
	//args := make([]reflect.Value, 0)
	fmt.Print(res)

}
