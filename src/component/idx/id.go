package idx

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

// spy 2022/1/10

func GetUUID() string {
	return uuid.New().String()
}

func GetShortUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func GetYYYYMMDD() string {
	return time.Now().Format("20060102")
}

func GetYYYYMMDDHH() string {
	return time.Now().Format("2006010215")
}

func GetYYYYMMDDHHMM() string {
	return time.Now().Format("200601021504")
}

func GetYYYYMMDDHHMMSS() string {
	return time.Now().Format("20060102150405")
}
