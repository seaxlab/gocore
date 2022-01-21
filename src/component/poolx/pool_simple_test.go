package poolx

import (
	"log"
	"runtime"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	numCPUs := runtime.NumCPU()

	pool := NewFunc(numCPUs, func(payload interface{}) interface{} {
		var result []byte

		// TODO: Something CPU heavy with payload
		time.Sleep(time.Duration(3) * time.Second)
		log.Printf("%v done\n", payload)

		return result
	})
	defer pool.Close()

	for i := 0; i < 100; i++ {
		// Funnel this work into our pool. This call is synchronous and will
		// block until the job is completed.
		go pool.Process(i)

		// 这里为了演示增加了go，增加go是为了模拟多个线程
	}

	time.Sleep(1 * time.Minute)
}
