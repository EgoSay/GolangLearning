package queue

type Queue []interface{} //可以存任意类型

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int)) //限定为int 类型
}
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return &head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
