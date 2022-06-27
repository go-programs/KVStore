package main

import (
	"fmt"
)

var store = map[string]string{
	"user_a":"Bangalore, Karnataka",
	"user_b":"Cochin, Kerala",
}

func get(key string) (string, error) {
	v, ok := store[key]
	if !ok {
		return v, fmt.Errorf("Value not found")
	}
	return v, nil
}

func put(key string, value string) {
	store[key] = value
	return
}

func del(key string) {
	delete(store, key)
	return
}

func main() {
	location_a, err := get("user_a")
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Location of user_a:",location_a)
	}
	del("user_a")
	fmt.Println("user_a details removed from key-value store")
	location_a, err = get("user_a")
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Location of user_a:",location_a)
	}
	put("user_c","Chennai, Tamilnadu")
	location_c, err := get("user_c")
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Location of user_c:",location_c)
	}
}