package queuex

import (
	"log"
	"testing"
	"time"
)

// spy 2022/1/13
func TestBlock(t *testing.T) {
	queue, _ := NewArrayBlockingQueue(2)

	go func() {
		for {
			time.Sleep(10 * time.Second)
			queue.Get() // Will unblock the first goroutine and add the last item
			log.Println("get----")
		}
	}()
	for i := 0; i < 20; i++ {
		log.Println("put---")
		queue.Put(1)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(3 * time.Minute)
}

func TestBlock2(t *testing.T) {
	queue, _ := NewArrayBlockingQueue(2)

	go func() {
		for {
			time.Sleep(10 * time.Second)
			queue.PopMulti(2) // Will unblock the first goroutine and add the last item
			log.Println("get----")
		}
	}()
	for i := 0; i < 20; i++ {
		log.Println("put---")
		queue.Put(1)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(3 * time.Minute)
}
