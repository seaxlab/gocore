package osx

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

// spy 2020/5/27

// Command creates a new exec.Cmd that will run with user privilege.
func Command(uid, command string, args ...string) (*exec.Cmd, error) {
	ucred, err := GetUserCredByUID(uid)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(command, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = ucred.Cred
	return cmd, nil
}

// Run creates and runs command with user privilege.
func Run(uid, command string, args ...string) error {
	cmd, err := Command(uid, command, args...)
	if err != nil {
		return err
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("command(su %s) %s: %v: %s",
			uid, cmd.Path, err, stderr.String())
	}
	return nil
}

// Output creates and runs command with user privilege and returns the output.
func Output(uid, command string, args ...string) ([]byte, error) {
	cmd, err := Command(uid, command, args...)
	if err != nil {
		return nil, err
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("command(su %s) %s: %v: %s (output: %s)",
			uid, cmd.Path, err, stderr.String(), string(out))
	}
	return out, nil
}

// CombinedOutput creates and runs command with user privilege and returns the
// combined output.
func CombinedOutput(uid, command string, args ...string) ([]byte, error) {
	cmd, err := Command(uid, command, args...)
	if err != nil {
		return nil, err
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("command(su %s) %s: %v: %s",
			uid, cmd.Path, err, string(out))
	}
	return out, nil
}

//kill single cmd
func KillSingle(pid int) (err error) {
	return syscall.Kill(pid, syscall.SIGKILL)
}

// 向进程组发信号
func KillGroup(pid int) (err error) {
	//，传递进程组PGID的时候要使用负数的形式
	return syscall.Kill(-pid, syscall.SIGKILL)
}
