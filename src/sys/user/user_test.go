package user

import (
	"log"
	"testing"
)

// spy 2020/5/27

func TestGetUserCred(t *testing.T) {
	print("root")
	print("shipengyan")
}

func print(username string) {
	userCred, _ := GetUserCred(username)

	log.Printf("uid=%s,gid=%s,username=%s,user.home=%s",
		userCred.User.Uid, userCred.User.Gid, userCred.User.Username, userCred.User.HomeDir)
}
