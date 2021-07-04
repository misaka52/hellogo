package queue

// A FIFO queue
type Queue []interface{}

// push a element into the queue
//
// 2
func (q *Queue) Push(value interface{}) {
	*q = append(*q, value)
}

/*
pop a element from the queue
2
3
	fdsa
*/
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// return if the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
