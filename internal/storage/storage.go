package storage

import "fmt"

var data = make(map[string]string)

func Set(key string, val string) {
	data[key] = val
}

func Get(key string) string {
	result, ok := data[key]
	if !ok {
		return ""
	}
	return result
}

func Debug() {
	fmt.Printf("[DEBUG] %v\n", data)
}

func Clear() {
	for k := range data {
		delete(data, k)
	}
}
