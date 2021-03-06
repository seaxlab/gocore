package clonex

import (
	"bytes"
	"encoding/gob"
)

// spy 2018/6/27

//实现结构之间的深拷贝

// Clone 深拷贝
// 注意：struct属性必须是公开（大写），可以拷贝数组和指针类型
func Clone(dst interface{}, src interface{}) error {
	return deepCopy1(dst, src)
}

//------------------------internal----------------------

func deepCopy1(dst interface{}, src interface{}) error {

	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

//
//func deepCopy2(dst interface{}, src interface{}) error {
//	//方式2, 相同field，不支持复杂类型，struct中嵌套数组字符串
//	return copier.Copy(dst, src)
//}
//
//func deepCopy3(dst interface{}, src interface{}) error {
//	bytes, err := json.Marshal(src)
//
//	if err != nil {
//		return err
//	}
//	json.Unmarshal(bytes, dst)
//	return nil
//}
