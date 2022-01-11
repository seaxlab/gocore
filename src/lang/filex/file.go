package filex

import (
	"fmt"
	string2 "github.com/seaxlab/gocore/src/lang/stringx"
	"io"
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
	return string2.Substr(p, 0, l)
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

// check path is file.
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
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
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
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		if sourceinfo, e := os.Stat(source); nil != e {
			err = os.Chmod(dest, sourceinfo.Mode())

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
