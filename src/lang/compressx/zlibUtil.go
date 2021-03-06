package compressx

import (
	"bytes"
	"compress/zlib"
	"io"
)

// spy 2020/1/21

func ZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

func ZlibUnCompress(src []byte) []byte {
	var out bytes.Buffer
	b := bytes.NewReader(src)
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}
