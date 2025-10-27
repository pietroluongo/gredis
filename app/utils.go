package main

import (
	"bytes"
)

func cleanupCommand(command []byte) string {
	return string(bytes.Trim(command, "\r\n\x00"))
}
