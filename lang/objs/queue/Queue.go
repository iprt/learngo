package queue

import (
	"github.com/pkg/errors"
)

// type 类型定义
// https://www.jianshu.com/p/a02cf41c0520
type Queue []int

func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

func (q *Queue) Pop() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head, nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
