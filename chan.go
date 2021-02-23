package main

import (
	"fmt"
	"sync"
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
	chanGroup.Wait()
	fmt.Println("x 的值：", x)
}

func main() {
	chanLock()
}
