package numberx

import (
	"log"
	"strconv"
	"testing"
)

// spy 2020/5/20

func TestName(t *testing.T) {
	v := "3.1415926535"
	v1, _ := strconv.ParseFloat(v, 32)
	v2, _ := strconv.ParseFloat(v, 64)

	log.Println(v1)
	log.Println(v2)
}

func TestRatio(t *testing.T) {

	log.Println(ratio(1, 6, 5))
}
