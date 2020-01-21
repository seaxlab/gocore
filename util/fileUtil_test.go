package util

import (
	"path/filepath"
	"testing"
)

func TestFileSize(t *testing.T) {

	userHome, err := UserHome()
	if err != nil {
		logger.Error("fail to get user home")
		return
	}
	fileSize := GetFileSize(filepath.Join(userHome, ".bash_history"))

	logger.Info("file size ", fileSize)
}
