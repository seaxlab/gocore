package timex

import (
	"fmt"
	"testing"
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
