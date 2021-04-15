package main

import (
	"context"
	"fmt"
	"time"
)

func loopGo(cxt context.Context, number int) {
	for {
		select {
		case val := <-cxt.Done():
			fmt.Printf("value %s %d", val, number)
			return
		}
	}
}

func testCancelContext() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i < 5; i++ {
		go loopGo(ctx, i)
	}

	time.Sleep(1 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
}

func main() {
	testCancelContext()
}
