package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Queue struct {
	elements []int
	front    int
	rear     int
	len      int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		elements: make([]int, capacity),
		front:    0,
		rear:     -1,
		len:      0,
	}
}

func (q *Queue) Enqueue(value int) bool {
	if q.len == len(q.elements) {
		return false
	}
	q.rear = (q.rear + 1) % (len(q.elements))
	q.elements[q.rear] = value
	q.len++
	return true
}

func (q *Queue) Dequeue() (int, bool) {
	if q.len == 0 {
		return 0, false
	}
	data := q.elements[q.front]
	q.front = (q.front + 1) % len(q.elements)
	q.len--
	return data, true
}

func main() {
	mu := sync.Mutex{}
	fullCond := sync.NewCond(&mu)
	emptyCond := sync.NewCond(&mu)
	queue := NewQueue(10)

	producer := func() {
		for {
			value := rand.Int()
			mu.Lock()
			for !queue.Enqueue(value) {
				fmt.Println("producer wait, queue is full..")
				fullCond.Wait()
			}
			mu.Unlock()
			emptyCond.Signal()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}
	}

	consumer := func() {
		for {
			mu.Lock()
			var v int
			for {
				var ok bool
				if v, ok = queue.Dequeue(); !ok {
					fmt.Println("Queue is empty")
					emptyCond.Wait()
					continue
				}
				break
			}
			mu.Unlock()
			fullCond.Signal()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			fmt.Println(v)
		}
	}

	for i := 0; i < 10; i++ {
		go producer()
	}

	for i := 0; i < 10; i++ {
		go consumer()
	}

	select {}

}
