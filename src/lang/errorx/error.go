package errorx

import (
	"errors"
	"fmt"
)

// spy 2022/1/12

func New(msg string) error {
	return errors.New(msg)
}

func NewF(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}
