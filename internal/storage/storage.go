package storage

import (
	"fmt"
	"sync"
)

type Storage struct {
	data map[string]string
}

var (
	instance *Storage
	once     sync.Once
)

func init() {
	once.Do(func() {
		fmt.Println("Creating storage instance...")
		instance = &Storage{
			data: make(map[string]string),
		}
	})

}

func Set(key string, value string) {
	instance.data[key] = value
}

func Get(key string) *string {
	val, ok := instance.data[key]
	if ok {
		return &val
	}
	return nil
}

func Debug() {
	fmt.Printf("%v\n", instance)
}
