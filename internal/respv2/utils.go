package respv2

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

const BYTES_TO_TRIM = "\n\r\x00"

func cleanInput(input []byte) []byte {
	return bytes.Trim(input, BYTES_TO_TRIM)
}

func extractArraySize(data []byte) (int, int, error) {
	// we're dealing with array - max size is 2^32 = 10 chars in dec
	arrSize := make([]byte, 10)
	i := 0
	for i < len(data) && unicode.IsNumber(rune(data[i])) {
		arrSize[i] = data[i]
		i += 1
	}
	convertedArrSize := strings.Trim(string(arrSize), "\x00")
	sz, err := strconv.Atoi(string(convertedArrSize))
	return sz, i, err
}
