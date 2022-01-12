package stringx

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// spy 2020/1/21

func TestToStr(t *testing.T) {
	assert.Equal(t, "1", ToStr(1))
	assert.Equal(t, "a", ToStr("a"))
}

func TestTrimSpace(t *testing.T) {
	fmt.Println(TrimSpaceString(" helloworld "))

	src := "  a\ta a \t  \n  \t\r  \r\n\n  b  b\t b \t"
	expect := "a\ta ab  b\t b"
	expectBytes := []byte(expect)

	if dst := TrimSpaceString(src); dst != expect {
		t.Errorf("TrimSpaceString(%q):\nhave %q\nwant %q\n", src, dst, expect)
	}
	if dst := TrimSpace([]byte(src)); !bytes.Equal(dst, expectBytes) {
		t.Errorf("TrimSpace(%q):\nhave %q\nwant %q\n", src, dst, expectBytes)
	}
}

func TestJoin(t *testing.T) {
	assert.Equal(t, JoinDefault("a", "b", "c"), "a:b:c")
	assert.Equal(t, Join("-", "a", "b", "c"), "a-b-c")
}
