package treex

import (
	"fmt"
	"github.com/seaxlab/gocore/src/lang/filex"
	"path/filepath"
	"testing"
	"time"
)

func TestRBTree(t *testing.T) {
	fmt.Println("consumer_config")

	filename, _ := filepath.Abs("../../../data/consumer_config.yml")
	words, _ := filex.GetContent(filename)
	fmt.Println("Total words: ", len(words))

	// Test RBTree
	startTime := time.Now()
	rb := NewRBTree()
	for _, word := range words {
		if rb.Contains(word) {
			rb.Set(word, rb.Get(word).(int)+1)
		} else {
			rb.Add(word, 1)
		}
	}
	for _, word := range words {
		rb.Contains(word)
	}
	fmt.Println("RB tree: ", time.Now().Sub(startTime))
	fmt.Println(rb.size)
}
