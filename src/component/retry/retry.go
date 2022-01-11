package retry

import (
	"math/rand"
	"time"
)

type Err error

type Stop struct {
	Err
}

func Retry(f func() error, attempts int, sleep time.Duration) error {
	if err := f(); err != nil {
		if s, ok := err.(Stop); ok {
			return s.Err
		}

		if attempts > 0 {
			delta := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + delta/2

			time.Sleep(sleep)
			attempts--
			return Retry(f, attempts, 2*sleep)
		}
		return err
	}
	return nil
}
