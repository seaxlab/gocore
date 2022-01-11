package filex

import (
	"fmt"
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

func TestWriterStr(t *testing.T) {
	userHome, _ := osx.UserHome()
	seaHome := filepath.Join(userHome, "sea", "test")
	testFile := filepath.Join(seaHome, "file.txt")

	err := CreateFile(testFile)
	fmt.Println(err)

	WriteString(testFile, "abcd\n")
}

func TestAppendStr(t *testing.T) {
	userHome, _ := osx.UserHome()
	seaHome := filepath.Join(userHome, "sea", "test")
	testFile := filepath.Join(seaHome, "file.txt")

	AppendString(testFile, "append it \n")
}

func TestGetContent(t *testing.T) {
	userHome, _ := osx.UserHome()
	seaHome := filepath.Join(userHome, "sea", "test")
	testFile := filepath.Join(seaHome, "file.txt")

	content, _ := GetContent(testFile)
	fmt.Println(content)
}
