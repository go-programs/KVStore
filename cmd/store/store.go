package store

import (
	"fmt"
	"sync"
)

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: map[string]string{}}

func Get(key string) (string, error) {
	store.RLock()
	v, ok := store.m[key]
	store.RUnlock()
	if !ok {
		return v, fmt.Errorf("Value not found")
	}
	return v, nil
}

func Put(key string, value string) {
	store.Lock()
	store.m[key] = value
	store.Unlock()
	return
}

func Del(key string) {
	store.Lock()
	delete(store.m, key)
	store.Unlock()
	return
}