package main

import (
	"fmt"
	"strings"
)

type Store struct {
	data map[string]string
}

func main() {
	// store := Store{data: make(map[string]string)}

	// store.set("first", "value")
	// value := store.get("first")

	// fmt.Println(value)
	input := "SET key value"
	words := strings.Fields(input)
	fmt.Println(words)
}

// func (s *Store) set(key, value string) {
// 	s.data[key] = value
// }

// func (s Store) get(key string) string {
// 	return s.data[key]
// }
