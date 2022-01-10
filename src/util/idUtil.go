package util

import (
	"github.com/google/uuid"
	"strings"
)

// spy 2022/1/10

func GetUUID() string {
	return uuid.New().String()
}

func GetShortUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
