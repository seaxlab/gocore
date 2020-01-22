package datetime

import "testing"

// spy 2020/1/22

func TestDatetimeFormat(t *testing.T) {
	logger.Info("now is ", Format(Now(), "YYYY-MM-DD hh:mm:ss"))
}

func TestParse2(t *testing.T) {
	time, err := Parse("2019-01-01", "YYYY-MM-DD")
	if err != nil {

	}
	logger.Info(time)
}
