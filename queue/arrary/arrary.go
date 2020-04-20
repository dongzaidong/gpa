package arrary

import (
	"errors"
	"sync"
)

type ArraryQ struct {
	items []interface{}

	head int
	tail int
	size int

	full bool
	len  int
	mu   sync.Mutex
}

func NewArraryQ(size int) *ArraryQ {
	return &ArraryQ{size: size, items: make([]interface{}, size)}
}

func (a *ArraryQ) Push(data interface{}) error {

	if a.size < 1 {
		return errors.New("[arrary] queue size is 0")
	}

	if a.full {
		return errors.New("[arrary] queue data is full")
	}

	a.items[a.tail] = data
	a.tail++
	a.len++
	if a.tail == a.size {
		a.tail = 0
	}
	if a.tail == a.head {
		a.full = true
	}

	return nil
}

func (a *ArraryQ) Pop() (interface{}, bool) {

	if a.head == a.tail && a.full == false {
		return nil, false
	}
	data := a.items[a.head]
	a.head++
	a.len--

	if a.head == a.size {
		a.head = 0
	}
	a.full = false
	return data, true

}
func (a *ArraryQ) Len() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.len
}

func (a *ArraryQ) IsEmpty() bool {

	a.mu.Lock()
	defer a.mu.Unlock()
	return a.len == 0
}
