package main

import (
	"bytes"
	"strings"
)

func cleanupCommand(command []byte) string {
	return strings.ToLower(string(bytes.Trim(command, "\r\n\x00")))
}
