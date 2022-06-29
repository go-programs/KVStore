package store

import (
	"fmt"
)

var store = map[string]string{}

func Get(key string) (string, error) {
	v, ok := store[key]
	if !ok {
		return v, fmt.Errorf("Value not found")
	}
	return v, nil
}

func Put(key string, value string) {
	store[key] = value
	return
}

func Del(key string) {
	delete(store, key)
	return
}