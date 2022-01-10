package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// spy 2020/4/15

func TestGetMatchValue(t *testing.T) {
	content := "123-ab,app=user,cc=cc-b"

	logger.Info(GetMatchValue(content, "app"))
	logger.Info(GetMatchValue(content, "cc"))
	assert.Equal(t, "user", GetMatchValue(content, "app"))
	assert.Equal(t, "cc-b", GetMatchValue(content, "cc"))
}
