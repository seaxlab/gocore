package filex

import (
	"fmt"
	"github.com/seaxlab/gocore/src/lang/osx"
	"log"
	"path/filepath"
	"testing"
)

var (
	userHome    string
	seaTestHome string
)

func init() {
	path, _ := osx.UserHome()
	userHome = path
	seaTestHome = filepath.Join(userHome, "sea", "test")
}

func TestGetParentDir(t *testing.T) {
	fmt.Println(GetParentDir(seaTestHome + "/file.txt"))
}

func TestGetFileName(t *testing.T) {
	fmt.Println(GetFileName(seaTestHome + "/file.txt"))
}

func TestFileSize(t *testing.T) {
	fileSize, _ := GetFileSize(filepath.Join(userHome, ".bash_history"))

	log.Printf("file size %v", fileSize)
}

func TestWriterStr(t *testing.T) {
	testFile := filepath.Join(seaTestHome, "file.txt")

	err := CreateFile(testFile)
	fmt.Println(err)

	WriteString(testFile, "abcd\n")
}

func TestAppendStr(t *testing.T) {
	testFile := filepath.Join(seaTestHome, "file.txt")

	AppendString(testFile, "append it \n")
}

func TestGetContent(t *testing.T) {
	testFile := filepath.Join(seaTestHome, "file.txt")

	content, _ := GetContent(testFile)
	fmt.Println(content)
}
