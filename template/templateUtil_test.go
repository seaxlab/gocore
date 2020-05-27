package template

import (
	"gitee.com/seaframework/go-core/log"
	"os"
	"testing"
)

// spy 2020/1/22
var logger = log.NewLogger(os.Stdout)

func TestTemplate(t *testing.T) {
	templateStr := "http://{host}/?q={query}&foo={bar}{bar}"

	varMap := map[string]interface{}{
		"host":  "www.baidu.com",
		"query": "helloworld",
		"bar":   "789",
	}

	content := Format(templateStr, varMap)
	logger.Info(content)
}
