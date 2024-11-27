/**
 * @Author codeAC
 * @Time: 2018/8/11 11:46
 * @Description
 */

package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

//如果没有缓冲区，channel发送没有接收者时就会出现死锁
func bufferedChannel() {
	c := make(chan int, 3) //缓冲区大小为3
	c <- 'a'
	c <- 'b'
	c <- 'c'
	go worker(1, c)
	time.Sleep(time.Millisecond)
}

//发送方才能close
func channelClose() {
	c := make(chan int)
	go worker(1, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
