package datetime

// spy 2020/1/21
import (
	"gitee.com/seaframework/go-core/log"
	"os"
	"testing"
)

// Logger
var logger = log.NewLogger(os.Stdout)

func TestFormat(t *testing.T) {
	logger.Info(FormatTime("2006-01-01"))
}
