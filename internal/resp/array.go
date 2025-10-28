package resp

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func arrayBuilder(data []byte) []RespMessage {
	sz, offset, err := extractArraySize(data)
	if err != nil {
		log.Error("Failed to extract array size")
		return []RespMessage{}
	}

	result := make([]RespMessage, sz)

	for i := 0; i < sz; i++ {
		data, err := internalParse(data[offset:])
		if err != nil {
			log.Error(fmt.Sprintf("failed to extract data %d - %s", i, err))
			continue
		}
		result[i] = *data
	}
	return result
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
