package urlx

import (
	"log"
	"testing"
)

// spy 2022/1/26

func TestBuildParam(t *testing.T) {
	params := make(map[string]string)

	params["code"] = "1"
	params["name"] = "3"

	log.Println(BuildParam(params))
}
