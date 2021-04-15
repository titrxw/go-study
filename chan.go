package main

import (
	"fmt"
	"sync"
	"time"
)

// 由于 x=x+1 不是原子操作
// 所以应避免多个协程对x进行操作
// 使用容量为1的信道可以达到锁的效果
func increment(group *sync.WaitGroup, ch chan bool, x *int) {
	defer group.Done()

	//信道的长度为1，当往信道中有值时，再次往信道中加值，会处于等待状态，当信道中为空时，处于等待中的协程才会继续往下执行
	ch <- true
	*x = *x + 1
	<-ch
}

func chanLock() {
	// 注意要设置容量为 1 的缓冲信道
	pipline := make(chan bool, 1)

	var chanGroup sync.WaitGroup
	chanGroup.Add(1000)

	var x int
	for i := 0; i < 1000; i++ {
		go increment(&chanGroup, pipline, &x)
	}
	fmt.Println("x 的值：", x)

	chanGroup.Wait()
}

func noBufferProductor(chanGroup *sync.WaitGroup, noBuffer chan int) {
	defer chanGroup.Done()
	noBuffer <- 1
	println("test")
}

func noBufferConsumer(chanGroup *sync.WaitGroup, noBuffer chan int) {
	defer chanGroup.Done()
	msg := <-noBuffer
	println(msg)
}

func chanNoBuffer() {
	var chanGroup1 sync.WaitGroup
	chanGroup1.Add(2)

	noBuffer := make(chan int)
	//无缓存信道，如果没有从信道中取数据的协程，会报错
	go noBufferConsumer(&chanGroup1, noBuffer)
	go noBufferProductor(&chanGroup1, noBuffer)

	chanGroup1.Wait()
}

func stopGo() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func main() {
	chanLock()

	chanNoBuffer()

	stopGo()
}
