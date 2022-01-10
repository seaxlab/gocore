package template

import (
	"github.com/seaxlab/gocore/log"
	"os"
	"strings"
	"testing"
	"time"
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

func TestName(t *testing.T) {
	now := time.Now().UTC()
	begin := now.Add(-1 * time.Hour * time.Duration(1))
	end := now.Add(time.Hour * time.Duration(1))
	logger.Info(begin.Format("2006-01-02T15:04:05.000Z"))
	logger.Info(end.Format("2006-01-02T15:04:05.000Z"))
}

func TestName2(t *testing.T) {
	region := "qingdaopro"
	logger.Info(strings.Contains(region, "qingdao"))
}
