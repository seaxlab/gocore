package timex

import (
	"github.com/metakeule/fmtdate"
	"time"
)

// spy 2020/1/22

func Now() time.Time {
	return time.Now()
}

// Format format datetime
// YYYY-MM-DD hh:mm:ss //注意这里是YYYY大写的
func Format(time time.Time, format string) string {
	return fmtdate.Format(format, time)
}

// Parse parse datetime
func Parse(datetimeStr string, format string) (time.Time, error) {
	return fmtdate.Parse(format, datetimeStr)
}
