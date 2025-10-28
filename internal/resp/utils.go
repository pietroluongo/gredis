package resp

import "bytes"

const BYTES_TO_TRIM = "\n\r\x00"

func cleanInput(input []byte) []byte {
	return bytes.Trim(input, BYTES_TO_TRIM)
}
