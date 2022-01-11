package filex

import (
	"fmt"
	"github.com/seaxlab/gocore/src/lang/stringx"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// RealPath get file real path
func RealPath(filePath string) string {
	p, err := filepath.Abs(filePath)
	if err != nil {
		log.Panicln("Get absolute path error.")
	}
	p = strings.Replace(p, "\\", "/", -1)
	l := strings.LastIndex(p, "/") + 1
	return stringx.Substr(p, 0, l)
}

// IsExists check path is exist
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// GetFileSize get the length in bytes of file of the specified path.
func GetFileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if nil != err {
		return -1, err
	}
	return fi.Size(), nil
}

// IsExist determines whether the file spcified by the given path is exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

// IsBinary determines whether the specified content is a binary file content.
func IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}

	return false
}

//IsFile check path is file.
func IsFile(path string) bool {
	b, err := os.Stat(path)
	if err != nil {
		return false
	}
	if b.IsDir() {
		return false
	}
	return true
}

// IsDir determines whether the specified path is a directory.
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if nil != err {
		fmt.Printf("Determines whether [%s] is a directory failed: [%v]", path, err)
		return false
	}

	return fio.IsDir()
}

// IsImg determines whether the specified extension is a image.
func IsImg(extension string) bool {
	ext := strings.ToLower(extension)

	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	default:
		return false
	}
}

// CreateFile 创建文件
func CreateFile(path string) error {
	// create parent dir if not exist
	parentPath := filepath.Dir(path)
	if _, err := os.Stat(parentPath); os.IsNotExist(err) {
		os.MkdirAll(parentPath, 0744)
	}

	// create file if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}

		defer file.Close()
	}
	return nil
}

// DeleteFile 删除文件
func DeleteFile(path string) error {
	_, err := os.Stat(path)

	// delete if exist
	if os.IsExist(err) {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}

// CopyFile copies the source file to the dest file.
func CopyFile(source string, dest string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		if sourceInfo, e := os.Stat(source); nil != e {
			err = os.Chmod(dest, sourceInfo.Mode())
			return
		}
	}

	return nil
}

// CopyDir copies the source directory to the dest directory.
func CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, err := os.Open(source)
	if err != nil {
		return err
	}

	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			return CopyDir(srcFilePath, destFilePath)
		} else {
			return CopyFile(srcFilePath, destFilePath)
		}
	}

	return nil
}

// GetContent read file content
func GetContent(filePath string) (string, error) {
	res, err := ioutil.ReadFile(filePath)
	return string(res), err
}

func WriteString(filePath string, content string) error {
	return Write(filePath, []byte(content))
}

func Write(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, 0644)
}

func AppendString(filePath string, content string) error {
	return Append(filePath, []byte(content))
}

func Append(filePath string, data []byte) error {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
