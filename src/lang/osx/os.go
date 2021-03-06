package osx

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// SeaHome get sea home
func SeaHome() (string, error) {
	user, err := user.Current()
	if err == nil {
		return filepath.Join(user.HomeDir, "sea"), nil
	}

	if IsWindows() {
		path, err := homeWindows()
		if err == nil {
			return filepath.Join(path, "sea"), nil
		}
	}

	path, err := homeUnix()
	if err == nil {
		return filepath.Join(path, "sea"), nil
	}
	return "", errors.New("unable to get sea home, plz check")
}

// UserHome returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func UserHome() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	if IsWindows() {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

// GetHostname get local hostname
func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

// IsWindows determines whether current OS is Windows.
func IsWindows() bool {
	return "windows" == runtime.GOOS
}

// GetWorkingDir get working dir
func GetWorkingDir() (string, error) {
	return os.Getwd()
}

// GetPwd gets the path of current working directory.
func GetPwd() string {
	file, _ := exec.LookPath(os.Args[0])
	pwd, _ := filepath.Abs(file)

	return filepath.Dir(pwd)
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
