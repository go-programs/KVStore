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

func main() {
	location_a, err := get("user_a")
	if(err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println("Location of user_a:",location_a)
}