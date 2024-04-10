package handlers

import "fmt"

func notFound(key string) error {
	return fmt.Errorf("key %s not found", key)
}
