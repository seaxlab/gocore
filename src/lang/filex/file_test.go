package filex

import (
	"github.com/seaxlab/gocore/src/lang/osx"
	"log"
	"path/filepath"
	"testing"
)

func TestFileSize(t *testing.T) {

	userHome, _ := osx.UserHome()
	fileSize, _ := GetFileSize(filepath.Join(userHome, ".bash_history"))

	log.Printf("file size %v", fileSize)
}
