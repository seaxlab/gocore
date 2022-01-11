package regexpx

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// spy 2020/4/15

func TestGetMatchValue(t *testing.T) {
	content := "123-ab,app=user,cc=cc-b"

	log.Println(GetMatchValue(content, "app"))
	log.Println(GetMatchValue(content, "cc"))
	assert.Equal(t, "user", GetMatchValue(content, "app"))
	assert.Equal(t, "cc-b", GetMatchValue(content, "cc"))
}

func TestRegExp(t *testing.T) {
	log.Println(IsLetter('a'))
	log.Println(IsLetter('1'))
	log.Println(IsChinaMobileString("17600001234"))
}
