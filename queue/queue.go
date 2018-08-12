package queue

//可以存任意类型
//type Queue []interface{}
type Queue []int

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int)) //限定为int 类型
}
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
