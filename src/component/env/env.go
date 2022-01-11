package env

import (
	"github.com/seaxlab/gocore/src/component/properties"
	"github.com/seaxlab/gocore/src/lang/filex"
	"github.com/seaxlab/gocore/src/lang/osx"
	"path/filepath"
)

// spy 2022/1/11

var securityProps properties.Properties

func GetSecurityValue(key string) string {
	if securityProps != nil {
		return securityProps[key]
	}
	return ""
}

func init() {
	path, _ := osx.UserHome()
	seaPath := filepath.Join(path, "sea/sea.password.properties")

	if filex.IsExist(seaPath) {
		props, err := properties.Read(seaPath)
		if err != nil {

			return
		}
		securityProps = props
	} else {
		securityProps = make(map[string]string)
	}

}
