package main

import (
	"Learning/queue"
	"fmt"
)

func main() {
	queues := queue.Queue{1}
	queues.Push(2)
	queues.Push(3)
	fmt.Println(queues.Pop())
	fmt.Println(queues.Pop())
	fmt.Println(queues.IsEmpty())
	fmt.Println(queues.Pop())
	fmt.Println(queues.IsEmpty())
	queues.Push("ss")
	fmt.Println(queues)
}
