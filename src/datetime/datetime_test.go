package datetime

import (
	"fmt"
	"testing"
)

// spy 2020/1/22

func TestDatetimeFormat(t *testing.T) {
	fmt.Printf("now=%s", Format(Now(), "YYYY-MM-DD hh:mm:ss"))
}

func TestParse2(t *testing.T) {
	time, err := Parse("2019-01-01", "YYYY-MM-DD")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time)
}
