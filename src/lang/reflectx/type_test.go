package reflectx

import (
	"log"
	"reflect"
	"testing"
)

// spy 2020/1/21

func TestIsBool(t *testing.T) {
	log.Println(IsBool("true"))
	log.Println(IsTrue("true"))
	log.Println(IsBool("true"))
	log.Println(IsBool(true))
}

func TestIsSlice(t *testing.T) {
	var a []int
	log.Println(IsSlice(a))
	var b [5]int
	log.Println(IsSlice(b))
}

func TestIsArray(t *testing.T) {
	var b [5]int
	log.Println(IsArray(b))
}

func TestName(t *testing.T) {
	a := 1
	//switch a.(type) {
	//case int:
	//
	//}
	log.Println(reflect.TypeOf(a).Kind())

	var b uint32
	b = 1
	log.Println(reflect.TypeOf(b).Kind() == reflect.Uint32)
}
