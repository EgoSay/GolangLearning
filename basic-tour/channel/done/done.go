/**
 * @Author codeAC
 * @Time: 2018/8/11 11:46
 * @Description
 */

package main

import (
	"fmt"
)

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		go func() { done <- true }()
	}
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
		//<- worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		//<- worker.done
	}
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo()
}
