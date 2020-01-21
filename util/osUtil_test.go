package util

import "testing"

// spy 2020/1/21

func TestGetHostname(t *testing.T) {
	logger.Info("hostname=", GetHostname())
}
