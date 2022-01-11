package env

import (
	"fmt"
	"testing"
)

// spy 2022/1/11

func TestGetSecurityValue(t *testing.T) {
	fmt.Println(GetSecurityValue("db_account_uat"))
}
