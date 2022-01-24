package treex

import (
	"fmt"
	"github.com/seaxlab/gocore/src/lang/filex"
	"path/filepath"
	"testing"
	"time"
)

func TestAVLTree(t *testing.T) {

	fmt.Println("consumer_config")

	filename, _ := filepath.Abs("../../../data/consumer_config.yml")
	words, _ := filex.GetContent(filename)
	fmt.Println("Total words: ", len(words))

	// Test AVL
	startTime := time.Now()
	avl := NewAVLTree()
	for _, word := range words {
		if avl.Contains(word) {
			avl.Set(word, avl.Get(word).(int)+1)
		} else {
			avl.Add(word, 1)
		}
	}
	for _, word := range words {
		avl.Contains(word)
	}
	fmt.Println("AVL: ", time.Now().Sub(startTime))
	fmt.Println(avl.size)
}
