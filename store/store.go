package store

import (
	"fmt"
	"sync"
)

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: map[string]string{}}

type Service interface {
	Get(string) (string, error)
	Put(string, string)
	Del(string)
}

type service struct {}

func NewService() Service {
	return &service{}
}

func (s *service) Get(key string) (string, error) {
	store.RLock()
	v, ok := store.m[key]
	store.RUnlock()
	if !ok {
		return v, fmt.Errorf("Value not found")
	}
	return v, nil
}

func (s *service) Put(key string, value string) {
	store.Lock()
	store.m[key] = value
	store.Unlock()
	return
}

func (s *service) Del(key string) {
	store.Lock()
	delete(store.m, key)
	store.Unlock()
	return
}