package clonex

import (
	"encoding/json"
	"github.com/seaxlab/gocore/src/lang/jsonx"
	"log"
	"reflect"
	"testing"
)

// spy 2018/6/27

type userModel struct {
	Name   string    `json:"name"`
	ID     int32     `json:"id"`
	ID2    int16     `json:"id2"`
	Roles  []string  `json:"roles"`
	Roles2 []*string `json:"roles2"`
}

type userRoleModel struct {
	Name   string    `json:"name"`
	ID     int64     `json:"id"`
	ID2    int8      `json:"id2"`
	Roles  []string  `json:"roles"`
	Roles2 []*string `json:"roles2"`
}

func TestDeepCopy(t *testing.T) {
	str1 := "1"
	str2 := "2"
	user1 := &userModel{Name: "smith", ID: 10, Roles: []string{"1", "2"}, Roles2: []*string{&str1, &str2}}
	user2 := &userModel{}

	log.Println(user1)
	log.Println(user2)

	Clone(user2, user1)

	log.Printf("after clone user1=%v\n", user1)
	log.Printf("after clone user2=%v\n", user2)

	log.Println("-----------")

	userRoleInstance := &userRoleModel{}

	Clone(userRoleInstance, user1)

	log.Printf("after clone user1=%v\n", user1)
	log.Printf("after clone userRoleInstance=%v\n", user2)
}

func TestDeepClone2(t *testing.T) {
	str1 := "1"
	str2 := "2"
	user1 := &userModel{Name: "smith", ID: 10, Roles: []string{"1", "2"}, Roles2: []*string{&str1, &str2}}
	user2 := &userModel{}

	log.Println(user1)
	log.Println(user2)

	origJSON, _ := json.Marshal(user1)
	log.Println(origJSON)

	str, _ := jsonx.ToJsonString(user1)
	log.Println(str)
	obj, _ := jsonx.DecodeJsonString(str)
	user3 := obj.(userModel)
	log.Println(user3)

	//json.Unmarshal(str, user2)

	log.Println(user1)
	log.Println(user2)
}

func TestDeepCopyLocal(t *testing.T) {

	var (
		str1 = "1"
		str2 = "2"
	)

	userInstance := &userModel{Name: "smith", ID: 10, Roles: []string{"1", "2"}, Roles2: []*string{&str1, &str2}}

	userRoleInstance := &userRoleModel{}

	log.Println(userInstance)
	log.Println(userRoleInstance)

	// 深拷贝
	StructCopy(userRoleInstance, userInstance)

	log.Println(userInstance)
	log.Println(userRoleInstance)

}

func DeepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}

func StructCopy(DstStructPtr interface{}, SrcStructPtr interface{}) {
	srcv := reflect.ValueOf(SrcStructPtr)
	dstv := reflect.ValueOf(DstStructPtr)
	srct := reflect.TypeOf(SrcStructPtr)
	dstt := reflect.TypeOf(DstStructPtr)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := DeepFields(reflect.ValueOf(SrcStructPtr).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}
