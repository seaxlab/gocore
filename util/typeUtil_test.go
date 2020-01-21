package util

import "testing"

// spy 2020/1/21

func TestIsBool(t *testing.T) {
	logger.Info(IsBool("true"))
	logger.Info(IsTrue("true"))
	logger.Info(IsBool("true"))
}
