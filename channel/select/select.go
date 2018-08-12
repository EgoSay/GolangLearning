/**
 * @Author codeAC
 * @Time: 2018/8/11 19:17
 * @Description
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(
				rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func createSelectWorker(id int) chan<- int {
	c := make(chan int)
	go selectWorker(id, c)
	return c
}

func selectWorker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createSelectWorker(0)
	var values []int
	tm := time.After(10 * time.Second) //设置程序结束时间
	tick := time.Tick(time.Second)
	var activeValue int
	for {
		var activeWorker chan<- int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
