package config

import (
	"fmt"
	"github.com/koding/multiconfig"
	"os"
	"strings"
	"sync"
)

// spy 2022/1/10
var (
	C    = new(Config)
	once sync.Once
)

//func init() {
//	MustLoad("")
//}

func MustLoad(filePaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, filePath := range filePaths {
			handled := false

			if strings.HasSuffix(filePath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: filePath})
				handled = true
			}
			if strings.HasSuffix(filePath, "conf") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: filePath})
				handled = true
			}
			if strings.HasSuffix(filePath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: filePath})
				handled = true
			}
			if strings.HasSuffix(filePath, "yaml") || strings.HasSuffix(filePath, "yml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: filePath})
				handled = true
			}

			if !handled {
				fmt.Println("config file invalid, valid file ext: .conf,.yaml,.yml,.toml,.json")
				os.Exit(1)
			}
		}

		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(C)

	})
}

type Config struct {
	Region string
	Env    string
}
