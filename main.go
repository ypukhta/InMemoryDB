package main

import (
	"fmt"
)

func main() {
	fmt.Println("In memory database demo")
	db := NewDB()
	key := "key1"
	val := "value1"
	fmt.Printf("Set: '%s:%s'\n", key, val)
	db.Set(key, val)
	val, _ = db.Get(key)
	fmt.Printf("Get: '%s:%s'\n", key, val)
	// Check tests for more cases
}
