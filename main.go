package main

import (
	"fmt"
	"reflect"

	"github.com/gpa/bus"
)

type Use struct {
}

func main() {
	ch := make(chan interface{})
	b := bus.NewEventBus()
	b.Subscribe("test", func(data interface{}) {
		ch <- data
	})

	b.Publish("test", 23333)

	d := <-ch
	fmt.Println("获取值", d)

	u := Use{}

	fmt.Println(reflect.TypeOf(u).Kind())
	fmt.Println(reflect.TypeOf(u).String())
}

func test() {

}
