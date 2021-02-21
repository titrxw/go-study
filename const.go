package main

import "fmt"

func main() {
	const num int = 12
	fmt.Println(num)

	const num1 = 13
	fmt.Println(num1)

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	fmt.Println(Unknown)
	fmt.Println(Female)
	fmt.Println(Male)

	const (
		Unknown1 = 2
		Female1
		Male1
	)
	fmt.Println(Female1)
	fmt.Println(Male1)

	const (
		test = iota
		test1
		test2
		test3 = 5 //如果中间改变了值 iota还会进行计数，但是下面数据的值不会自增
		test4
		test6
		test5 = iota
	)
	fmt.Println(test)
	fmt.Println(test1)
	fmt.Println(test1)
	fmt.Println(test4)
	fmt.Println(test6)
	fmt.Println(test5)
	//0 1 2 5 5 6

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
