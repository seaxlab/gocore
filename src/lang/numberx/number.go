package numberx

import (
	"fmt"
	"strconv"
)

// spy 2020/5/20

// IntToStr int到string
func IntToStr(value int) string {
	return strconv.Itoa(value)
}

// Int64ToStr int64到string
func Int64ToStr(value int64) string {
	return strconv.FormatInt(value, 10)
}

// Float32ToStr float32 to str
func Float32ToStr(value float64) string {
	return strconv.FormatFloat(value, 'E', -1, 32)
}

// Float64ToStr float64 to str
func Float64ToStr(value float64) string {
	return strconv.FormatFloat(value, 'E', -1, 64)
}

// string to int
func strToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

// string to int64
func strToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

// str to float32
func strToFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 32)
}

// str to float 64
func strToFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

//string := strconv.FormatFloat(float32, 'E', -1, 32)
//string := strconv.FormatFloat(float64, 'E', -1, 64)
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)

// 这里会产生精度问题

func ratio4(part, total int) (float64, error) {
	return ratio(part, total, 4)
}

func ratio(part, total, scale int) (float64, error) {
	if scale < 0 {
		scale = 4
	}
	// such as %.2f
	format := fmt.Sprintf("%s.%df", "%", scale)

	// 这里会四舍五入
	return strconv.ParseFloat(fmt.Sprintf(format, float64(part)/float64(total)), 64) // 保留[scale]位小数
}
