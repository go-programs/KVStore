package store

import (
	"fmt"
)

var store = map[string]string{
	"user_a":"Bangalore, Karnataka",
	"user_b":"Cochin, Kerala",
}

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

func main() {
	location_a, err := Get("user_a")
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Location of user_a:",location_a)
	}
	Del("user_a")
	fmt.Println("user_a details removed from key-value store")
	location_a, err = Get("user_a")
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Location of user_a:",location_a)
	}
	Put("user_c","Chennai, Tamilnadu")
	location_c, err := Get("user_c")
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Location of user_c:",location_c)
	}
}