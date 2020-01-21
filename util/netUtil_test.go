package util

import "testing"

// spy 2020/1/21

func TestLocalIP(t *testing.T) {
	logger.Info(GetLocalIP())
}
