package queuex

import (
	"fmt"
	list "github.com/seaxlab/gocore/src/lang/listx"
	"github.com/seaxlab/gocore/src/lang/numberx"
	"testing"
	"time"
)

// spy 2022/1/11
type Event struct {
	OrderNo string `json:"orderNo"`
}

func TestQueue(t *testing.T) {
	eventQueue := list.NewSafeListLimited(10000000)

	go func() {
		for {
			events := eventQueue.PopBackBy(1)
			if len(events) == 0 {
				time.Sleep(3 * time.Second)
				continue
			}
			for _, el := range events {
				event := el.(Event)
				fmt.Printf("orderNo=%s\n", event.OrderNo)
			}
		}
	}()

	for i := 0; i < 100; i++ {
		eventQueue.PushFront(Event{OrderNo: numberx.IntToStr(i)})
		time.Sleep(2 * time.Second)
	}
	time.Sleep(1 * time.Minute)
}
