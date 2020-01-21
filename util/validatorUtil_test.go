package util

import "testing"

// spy 2020/1/21

func TestValidatorUtil(t *testing.T) {
	logger.Info(IsLetter('a'))
	logger.Info(IsLetter('1'))
	logger.Info(IsChinaMobileString("17600001234"))

}
