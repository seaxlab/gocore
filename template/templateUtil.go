package template

import (
	"github.com/valyala/fasttemplate"
)

// spy 2020/1/22

func Format(template string, varMap map[string]interface{}) string {
	t := fasttemplate.New(template, "{", "}")
	content := t.ExecuteString(varMap)

	return content
}

//TODO 自定义，目前支持{name}占位符
