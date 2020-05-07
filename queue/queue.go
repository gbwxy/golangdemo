package queue

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}
func (q *Queue) Pop() int {
	if q == nil {
		panic("error.")
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
