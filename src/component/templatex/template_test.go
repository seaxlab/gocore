package templatex

import (
	"fmt"
	"testing"
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
