package config

import (
	"fmt"
	"github.com/seaxlab/gocore/src/lang/osx"
	"os"
	"testing"
)

// spy 2022/1/11

func TestConfig(t *testing.T) {
	path, _ := os.Getwd()
	fmt.Println(path)

	p, _ := osx.GetWorkingDir()
	fmt.Println(p)

	filePaths := []string{
		path + "/../../etc/config-test.local.yml",
	}
	MustLoad(filePaths...)
	MustLoad(path + "/../../etc/config-test2.local.yml")

	fmt.Println(C.Env)
	fmt.Println(C.Region)
}
