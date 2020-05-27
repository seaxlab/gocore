package user

import (
	"fmt"
	"testing"
)

// spy 2020/5/27

func TestGetUserCred(t *testing.T) {
	fmt.Println(GetUserCred("shipengyan"))
}
