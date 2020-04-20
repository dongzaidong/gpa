package bus

import (
	"reflect"
	"sync"
)

type PubReq struct {
	Topic string
	Param []reflect.Value
}

type EventBus struct {
	notice chan *PubReq
	done   chan int
	once   sync.Once
	mu     sync.Mutex
	topics map[string][]reflect.Value
}

func NewEventBus() *EventBus {
	e := &EventBus{
		notice: make(chan *PubReq),
		done:   make(chan int),
		topics: make(map[string][]reflect.Value),
	}
	go e.Run()
	return e

}
func (e *EventBus) Publish(topic string, args ...interface{}) {
	params := make([]reflect.Value, 0)

	for _, arg := range args {
		params = append(params, reflect.ValueOf(arg))
	}
	req := &PubReq{
		Topic: topic,
		Param: params,
	}

	e.notice <- req
}

func (e *EventBus) Subscribe(topic string, fn interface{}) {
	if topic != "" && reflect.TypeOf(fn).Kind() == reflect.Func {
		callback := reflect.ValueOf(fn)

		e.mu.Lock()
		defer e.mu.Unlock()
		e.topics[topic] = append(e.topics[topic], callback)

	}
}

func (e *EventBus) Unsubscribe(topic string, fn interface{}) {
	if topic != "" && reflect.TypeOf(fn).Kind() == reflect.Func {
		callback := reflect.ValueOf(fn)

		e.mu.Lock()
		defer e.mu.Unlock()
		var subs []reflect.Value

		for _, s := range e.topics[topic] {
			if s != callback {
				subs = append(subs, s)
			}
		}
		e.topics[topic] = subs

	}

}

func (e *EventBus) Stop() {
	e.once.Do(func() {
		close(e.done)
	})
}
func (e *EventBus) Run() {
	for {
		select {
		case <-e.done:
			return
		case info := <-e.notice:
			e.mu.Lock()
			callbacks, found := e.topics[info.Topic]
			e.mu.Unlock()

			if !found {
				continue
			}

			for _, cb := range callbacks {
				go e.execute(cb, info.Param)
			}

		}
	}
}

func (e *EventBus) execute(callback reflect.Value, args []reflect.Value) {
	callback.Call(args)
}
