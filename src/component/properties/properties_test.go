package properties

import (
	"fmt"
	"github.com/seaxlab/gocore/src/lang/osx"
	"path/filepath"
	"testing"
)

// spy 2022/1/11

func TestReadProperties(t *testing.T) {
	pwd, _ := osx.GetWorkingDir()
	path := filepath.Join(pwd, "../../../etc/test.properties")

	properties, err := Read(path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(properties)
	fmt.Printf("a=%s\n", properties["a"])
}
