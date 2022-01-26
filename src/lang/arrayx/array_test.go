package arrayx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// spy 2020/5/1

func TestToStrArray(t *testing.T) {
	at := make(map[string]interface{})
	at["atMobiles"] = []string{"123", "123"}
	at["atMobiles"] = append(at["atMobiles"].([]string), []string{"333"}...)

	fmt.Println(at["atMobiles"])
}

func TestAppend(t *testing.T) {
	myStringSlice := []string{"first", "second", "third"}

	myStringSlice = append(myStringSlice, []string{"fourth", "fift"}...)
	myStringSlice = append(myStringSlice, "sixth", "seventh")

	fmt.Println(myStringSlice)
}

func TestShuffle(t *testing.T) {
	data := []interface{}{"first", "second", "third", "111", "23"}

	Shuffle(data)
	fmt.Println(data)
}

func TestShuffleInt(t *testing.T) {
	data := []int{5, 4, 6, 7, 9, 10}

	ShuffleInt(data)
	fmt.Println(data)
}

func TestShuffleIntNew(t *testing.T) {
	data := []int{5, 4, 6, 7, 9, 10}
	newData := ShuffleIntNew(data)
	fmt.Println(data)
	fmt.Println(newData)
}

func MyMapFunc(v interface{}) interface{} {
	return v.(int) * 3
}

func TestMap(t *testing.T) {

	var myArray = []interface{}{1, 2, 3, 4}
	result := Map(myArray, MyMapFunc) // [3, 6, 9, 12]

	fmt.Println(result)
}

func TestCompact(t *testing.T) {
	var arr = []interface{}{1, 2, 3, 4, nil, 5, "hello world", nil}

	result := Compact(arr)
	assert.Equal(t, 6, len(result))
	assert.Equal(t, 1, result[0].(int))
	assert.Equal(t, 2, result[1].(int))
	assert.Equal(t, 3, result[2].(int))
	assert.Equal(t, 4, result[3].(int))
	assert.Equal(t, 5, result[4].(int))
	assert.Equal(t, "hello world", result[5].(string))
}

func TestFlatten(t *testing.T) {
	arr1 := []interface{}{1, 2, 3, 4}    // [1, 2, 3, 4]
	arr2 := []interface{}{5, 6, 7, arr1} // [5, 6, 7, [1, 2, 3, 4]]
	result := Flatten(arr2)              // [5, 6, 7, 1, 2, 3, 4]

	log.Println(result)
}

func TestMerge(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}

	z := append(x, y...)
	fmt.Println(z)
}

// 还不如上面的简洁
func TestMerge2(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}

	// []T to an []interface{}
	// https://go.dev/doc/faq#convert_slice_of_interface
	x2 := make([]interface{}, len(x))
	for i, v := range x {
		x2[i] = v
	}

	y2 := make([]interface{}, len(y))
	for i, v := range y {
		y2[i] = v
	}

	a := Merge(x2, y2)
	fmt.Println(a)
}

func TestMergeString(t *testing.T) {
	x := []string{"a", "b"}
	y := []string{"c"}

	result := MergeString(x, y)
	fmt.Println(result)
}

func TestUnique(t *testing.T) {
	var myArray = []interface{}{1, 2, 3, 3}
	result := Unique(myArray)
	assert.Equal(t, 3, len(result))
	assert.Equal(t, 1, result[0].(int))
	assert.Equal(t, 2, result[1].(int))
	assert.Equal(t, 3, result[2].(int))

	var myArray2 = []interface{}{"something", 1, 1.0, "you know what", 1, "something"}
	result = Unique(myArray2)
	assert.Equal(t, 4, len(result))
	assert.Equal(t, "something", result[0].(string))
	assert.Equal(t, 1, result[1].(int))
	assert.Equal(t, 1.0, result[2].(float64))
	assert.Equal(t, "you know what", result[3].(string))

	assert.Nil(t, Unique(nil))

	var myArray3 = []interface{}{"lone"}
	assert.Equal(t, myArray3, Unique([]interface{}(myArray3)))

	var myArray4 = []interface{}{"lone", "lone"}
	result = Unique(myArray4)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "lone", myArray4[0].(string))
}
