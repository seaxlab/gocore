package arrayx

import (
	"fmt"
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
	data := []int{5, 4, 6, 7, 9, 10}

	Shuffle(data)
	fmt.Println(data)
}
