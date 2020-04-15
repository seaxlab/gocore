package mytest

import (
	"fmt"
	"log"
	"regexp"
	"testing"
	"time"
)

// spy 2020/4/11

func TestName(t *testing.T) {
	str := "12cc=1,app=user-22,acb"

	//reg:= regexp.MustCompile("(app)=([0-9a-zA-Z\\-]+)")

	reg := regexp.MustCompile("(app)=([^,]*)")
	one := reg.FindAllStringSubmatch(str, -1)

	fmt.Println(one) //[[app=user-22 app user-22]]

	if one == nil || len(one) == 0 {
		return
	}
	if len(one[0]) == 3 {
		fmt.Println(one[0][2])
	}

}

func TestDate(t *testing.T) {

	now := time.Now()
	today := now.Format("2006-01-02")

	h := fmt.Sprintf("-%dh", 24) //减去24小时（前一天）
	dh, _ := time.ParseDuration(h)
	yesterday := now.Add(dh).Format("2006-01-02") //系统前一天日期格式转换

	log.Println(today)
	log.Println(yesterday)

	fmt.Println(today)
	fmt.Println(yesterday)
}
