package queue

//Queue A FIFO queue
type Queue []interface{}

//Push the element into the queue
//      e.g. q.push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)

}

//Pop a element from the queue
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//IsEmpty Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
