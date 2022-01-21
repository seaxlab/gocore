package stringx

// spy 2020/1/21

import (
	"bytes"
	"fmt"
	"strings"
)

var newlineBytes = []byte{'\n'}

func ToStr(a interface{}) string {
	return fmt.Sprint(a)
}

// IsEmpty returns true if the string is empty
func IsEmpty(text string) bool {
	return len(text) == 0
}

// IsNotEmpty returns true if the string is not empty
func IsNotEmpty(text string) bool {
	return !IsEmpty(text)
}

// IsBlank returns true if the string is blank (all whitespace)
func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

// IsNotBlank returns true if the string is not blank
func IsNotBlank(text string) bool {
	return !IsBlank(text)
}

func ContainsStr(s, substr string) bool {
	return strings.Contains(s, substr)
}

func StartsWith(s, substr string) bool {
	if substr != "" && Substr(s, 0, len([]rune(substr))) == substr {
		return true
	}
	return false
}

func EndsWith(s, substr string) bool {
	if Substr(s, -len([]rune(substr)), len(s)) == substr {
		return true
	}
	return false
}

//TrimSpace 去掉 src 开头和结尾的空白, 如果 src 包括换行, 去掉换行和这个换行符两边的空白
//  NOTE: 根据 '\n' 来分行的, 某些系统或软件用 '\r' 来分行, 则不能正常工作.
func TrimSpace(src []byte) []byte {
	bytesArr := bytes.Split(src, newlineBytes)
	for i := 0; i < len(bytesArr); i++ {
		bytesArr[i] = bytes.TrimSpace(bytesArr[i])
	}
	return bytes.Join(bytesArr, nil)
}

//TrimSpaceString 去掉 src 开头和结尾的空白, 如果 src 包括换行, 去掉换行和这个换行符两边的空白
//  NOTE: 根据 '\n' 来分行的, 某些系统或软件用 '\r' 来分行, 则不能正常工作.
func TrimSpaceString(src string) string {
	strs := strings.Split(src, "\n")
	for i := 0; i < len(strs); i++ {
		strs[i] = strings.TrimSpace(strs[i])
	}
	return strings.Join(strs, "")
}

// string to []bytes
func ToBytes(src string) []byte {
	return []byte(src)
}

// Substring 开始位置+结束位置
func Substring(str string, start, end int) string {
	if end > start {
		return ""
	}
	charlist := []rune(str)
	l := len(charlist)

	if start > l {
		start = l
	}

	if end > l {
		end = l
	}

	return string(charlist[start:end])

}

// Substr 开始位置+字符个数
func Substr(s string, start int, strLength int) string {
	charlist := []rune(s)
	l := len(charlist)
	length := 0
	end := 0

	if strLength <= 0 {
		length = l
	} else {
		length = strLength
	}

	if start < 0 {
		start = l + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > l {
		start = l
	}

	if end < 0 {
		end = 0
	}

	if end > l {
		end = l
	}

	return string(charlist[start:end])
}

// JoinDefault 默认通过“:”组装
func JoinDefault(args ...string) string {
	return Join(":", args...)
}

// Join 通过split组装
func Join(split string, args ...string) string {
	if len(args) == 0 {
		panic("args参数不能为空")
	}

	var buffer bytes.Buffer
	for i, v := range args {
		if i != 0 {
			buffer.WriteString(split)
		}
		buffer.WriteString(v)
	}
	return buffer.String()
}
