/**
 * @Author codeAC
 * @Time: 2018/8/11 11:46
 * @Description
 */

package main

import (
	"fmt"
	"sync"
)

func createWorkerWithWaitGroup(id int, wg *sync.WaitGroup) workerWithWaitGroup {
	w := workerWithWaitGroup{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorkWithWaitGroup(id, w.in, w.done)
	return w
}

type workerWithWaitGroup struct {
	in   chan int
	done func()
}

func doWorkWithWaitGroup(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func chanDemoWithWaitGroup() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	var wg sync.WaitGroup
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemoWithWaitGroup()
}
