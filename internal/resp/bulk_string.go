package resp

import (
	"fmt"
)

func bulkStringBuilder(data []byte) (*RespMessage, int, error) {
	log.Info(fmt.Sprintf("Parsing bulk string %s", string(data)))
	sz, offset, err := extractArraySize(data)
	if err != nil {
		log.Error("Failed to extract bulk string size")
		return nil, 0, err
	}
	log.Info(fmt.Sprintf("bulk string size is %d", sz))

	msg := data[offset+2 : offset+2+sz]

	return &RespMessage{Kind: BulkString, Content: string(msg)}, offset + 2 + 2 + sz + 1, nil
}
