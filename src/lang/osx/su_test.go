package osx

import (
	"log"
	"testing"
)

// spy 2020/5/28

func TestName(t *testing.T) {
	userCred, _ := GetUserCred("shipengyan")

	cmd, err := Command(userCred.User.Uid, "sh", "-c", "\"ls -l\"")
	err = cmd.Start()

	if err != nil {
		log.Printf("can't start command: %v", err)
	}
}
