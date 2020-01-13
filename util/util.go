package util

import (
	"../log"
	"os"
)

// Logger
var logger = log.NewLogger(os.Stdout)
