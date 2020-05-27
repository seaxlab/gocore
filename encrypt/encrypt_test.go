package encrypt

import (
	"gitee.com/seaframework/go-core/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// spy 2020/1/21

// Logger
var logger = log.NewLogger(os.Stdout)

func TestEncrypt(t *testing.T) {
	assert.Equal(t, Md5("1234561"), "e10adc3949ba59abbe56e057f20f883e")
}
