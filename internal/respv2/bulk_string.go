package respv2

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/internal/logger"
)

type BulkStringNode struct {
	Data string
}

func (BulkStringNode) node() {}

func bulkStringBuilder(data []byte) (Node, int, error) {
	logger.Log.Info(fmt.Sprintf("Parsing bulk string %s", string(data)))
	bulkStringSize, offset, err := extractArraySize(data)
	fmt.Printf("%d %d\n", bulkStringSize, offset)
	if err != nil {
		logger.Log.Error("Failed to extract bulk string size")
		return nil, 0, err
	}
	logger.Log.Info(fmt.Sprintf("bulk string size is %d", bulkStringSize))

	msg := data[offset+2 : offset+2+bulkStringSize]

	fmt.Printf("msg = %s\n", msg)

	return BulkStringNode{Data: string(msg)}, offset + 2 + 2 + bulkStringSize + 1, nil
}
