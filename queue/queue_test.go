/**
 * @Author codeAC
 * @Time: 2018/8/10 19:48
 * @Description
 */
package queue

import "fmt"

func ExampleQueue_Pop() {
	queues := Queue{1}
	queues.Push(2)
	fmt.Println(queues.Pop())
	fmt.Println(queues.IsEmpty())
	fmt.Println(queues.Pop())
	fmt.Println(queues.IsEmpty())
	//Output:
	// 1
	// false
	// 2
	// true
}
