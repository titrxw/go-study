package main

import (
	"fmt"
	"time"
)

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}

func notDefault() {
	c1 := make(chan byte, 3)
	c2 := make(chan string, 1)
	timeout := make(chan bool, 1)

	//往通道中添加3个值，最多三个
	c1 <- 'a'
	c1 <- 255
	c1 <- 251
	//c1 <- 255

	test := <-c1
	test1 := <-c1
	test2 := <-c1
	println(test)
	println(test1)
	println(test2)
	//到这里通道c2的值全部释放

	c2 <- "df"

	go makeTimeout(timeout, 2)

	println(1)

	//在运行 select 时，会遍历所有（如果有机会的话）的 case 表达式，只要有一个信道有接收到数据，那么 select 就结束
	//所以会从c2通道中获取到值，输出 c2 received df
	//如果不向c2通道中填充数据，只能等sleep结束后，执行timeout通道的填充，这时输出timeout exit
	//如果你没有写 default 分支，select 就会阻塞，直到有某个 case 可以命中
	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	case <-timeout:
		fmt.Println("Timeout, exit.")
	}
}

func hasDefault() {
	chanTest := make(chan string, 1)
	//如果通道中获取不到数据，执行default分支
	select {
	case <-chanTest:
		print("no default")
	default:
		print("default")
	}

	chanTest <- "re"
	select {
	case msg := <-chanTest:
		print(msg)
	default:
		print("default")
	}

	//表达式向通道中添加数据结果为真，所以输出test
	select {
	case chanTest <- "test":
		print("test")
	default:
		print("default")
	}
}

func closeChan(closech chan int) {
	select {
	case val := <-closech:
		fmt.Printf("close %d", val)
	}
}

func triggerCloseChan() {
	closech := make(chan int)
	go closeChan(closech)
	time.Sleep(1 * time.Second)
	close(closech)
	time.Sleep(1 * time.Second)
}

func random() {
	chanTest := make(chan string, 1)
	chanTest1 := make(chan int, 1)

	chanTest <- "test"
	chanTest1 <- 1

	//以下两个case都满足条件， 执行结果有随机性,执行结果可能是输出test，也可能是输出1
	select {
	case msg := <-chanTest:
		print(msg)
	case msg := <-chanTest1:
		print(msg)
	default:
		print("default")
	}
}

func main() {
	//仅能用于 信道/通道
	//跟 switch-case 很像
	notDefault()
	hasDefault()

	random()
	random()

	triggerCloseChan()
}
