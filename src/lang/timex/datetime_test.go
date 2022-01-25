package timex

import (
	"fmt"
	"testing"
	"time"
)

// spy 2020/1/22

func TestDatetimeFormat(t *testing.T) {
	fmt.Printf("now=%s", Format(Now(), FORMAT_YYYY_MM_DD_HH_MM_SS))
}

func TestParse2(t *testing.T) {
	time, err := Parse("2019-01-01", FORMAT_YYYY_MM_DD)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time)
}

func TestName(t *testing.T) {
	now := time.Now().UTC()
	begin := now.Add(-1 * time.Hour * time.Duration(1))
	end := now.Add(time.Hour * time.Duration(1))
	fmt.Println(begin.Format("2006-01-02T15:04:05.000Z"))
	fmt.Println(end.Format("2006-01-02T15:04:05.000Z"))
}
