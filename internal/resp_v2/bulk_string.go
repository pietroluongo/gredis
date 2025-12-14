package respv2

import (
	"fmt"
)

func bulkStringBuilder(data []byte) (Node, int, error) {
	log.Info(fmt.Sprintf("Parsing bulk string %s", string(data)))
	bulkStringSize, offset, err := extractArraySize(data)
	fmt.Printf("%d %d\n", bulkStringSize, offset)
	if err != nil {
		log.Error("Failed to extract bulk string size")
		return nil, 0, err
	}
	log.Info(fmt.Sprintf("bulk string size is %d", bulkStringSize))

	msg := data[offset+2 : offset+2+bulkStringSize]

	fmt.Printf("msg = %s\n", msg)

	return BulkStringNode{Data: string(msg)}, offset + 2 + 2 + bulkStringSize + 1, nil
}
