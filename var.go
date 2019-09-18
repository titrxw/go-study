package go_study

import (
	"fmt"
	"os"
)

/**
全局变量
:=只能用于局部变量, 不能用于全局变量
只要是通过该方式初始化的变量，都是局部变量，就算变量是全局变量， 再该作用范围内，该变量也是局部变量
*/
var height, country = 12, "zhongg"

func global1() {
	height := 1
	fmt.Println(height)
}

func global2() {
	height = 4
	fmt.Println(height)
}

func begin() {
	/**
	也可以先声明变量再进行初始化
	*/
	var time int
	time = 1
	fmt.Println(time)

	/**
	多个变量也可以先声明变量再进行初始化
	*/
	var name, body string
	name = "123"
	body = "456"
	name, body = "123", "456"
	fmt.Printf(name)
	fmt.Printf(body)
}

func person() {
	var num int = 1
	fmt.Printf(string(num))

	/**
	系统自动变量类型
	*/
	var age = 1
	fmt.Println(age)

	/**
	不写var进行变量的初始化，包括变量的类型
	*/
	name := "sdfsdfsdfsdf"
	fmt.Printf(name)

	/**
	不同的变量类型同时初始化，以及变量的类型
	*/
	var sex, city = 12, "tai yuan"
	fmt.Println(sex)
	fmt.Printf(city)

	/**
	这种写法只能在函数体内使用
	*/
	hair, body := "blue", "fat"
	fmt.Printf(hair)
	fmt.Printf(body)

	fmt.Printf(country)
	fmt.Println(height)
}

var cwd string = "test"

func test1() {
	cwd, err := os.Getwd() // NOTE: wrong!
	fmt.Print(cwd)
	fmt.Print(err)
}

func main() {
	test1()
	global1()
	fmt.Println(height)
	global2()
	fmt.Println(height)

	begin()
	person()

	var num = 4
	var num1 *int = &num
	*num1 = 5
	fmt.Println(num)
}
