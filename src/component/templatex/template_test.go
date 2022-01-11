package templatex

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// spy 2020/1/22

func TestTemplate(t *testing.T) {
	templateStr := "http://{host}/?q={query}&foo={bar}{bar}"

	varMap := map[string]interface{}{
		"host":  "www.baidu.com",
		"query": "helloworld",
		"bar":   "789",
	}

	content := Format(templateStr, varMap)
	fmt.Println(content)
}

func TestName(t *testing.T) {
	now := time.Now().UTC()
	begin := now.Add(-1 * time.Hour * time.Duration(1))
	end := now.Add(time.Hour * time.Duration(1))
	fmt.Println(begin.Format("2006-01-02T15:04:05.000Z"))
	fmt.Println(end.Format("2006-01-02T15:04:05.000Z"))
}

func TestName2(t *testing.T) {
	region := "qingdaopro"
	fmt.Println(strings.Contains(region, "qingdao"))
}
