package test

import (
	"github.com/seaxlab/gocore/src/lang/timex"
	"log"
	"testing"
	"time"
)

// spy 2022/1/11
func TestTime(t *testing.T) {
	log.Println(timex.Format(time.Now(), timex.FORMAT_YYYY_MM_DD))
}
