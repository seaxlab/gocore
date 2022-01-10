package qrcode

import (
	"github.com/seaxlab/gocore/src/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

// spy 20/1/27

func TestEncode(t *testing.T) {
	userHome, _ := util.UserHome()

	EncodeSimple("http://www.baidu.com", userHome+"/test/qr.png")
}

func TestDecode(t *testing.T) {
	userHome, _ := util.UserHome()

	content, _ := Decode(userHome + "/test/qr.png")
	assert.Equal(t, content, "http://www.baidu.com")
}
