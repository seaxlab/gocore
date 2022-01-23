package setx

import (
	"sync"
)

// SafeStringSet 线程安全的Set
type SafeStringSet struct {
	sync.RWMutex
	M map[string]bool
}

func NewSafeStringSet() *SafeStringSet {
	return &SafeStringSet{
		M: make(map[string]bool),
	}
}

func (this *SafeStringSet) Add(key string) {
	this.Lock()
	this.M[key] = true
	this.Unlock()
}

func (this *SafeStringSet) Remove(key string) {
	this.Lock()
	delete(this.M, key)
	this.Unlock()
}

func (this *SafeStringSet) Clear() {
	this.Lock()
	this.M = make(map[string]bool)
	this.Unlock()
}

func (this *SafeStringSet) Contains(key string) bool {
	this.RLock()
	_, exists := this.M[key]
	this.RUnlock()
	return exists
}

func (this *SafeStringSet) Size() int {
	this.RLock()
	len := len(this.M)
	this.RUnlock()
	return len
}

func (this *SafeStringSet) ToSlice() []string {
	this.RLock()
	defer this.RUnlock()

	count := len(this.M)
	if count == 0 {
		return []string{}
	}

	r := []string{}
	for key := range this.M {
		r = append(r, key)
	}

	return r
}
