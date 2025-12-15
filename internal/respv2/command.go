package respv2

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/internal/logger"
)

type CommandNode struct {
	Command []BulkStringNode // enable support for ACL commands, etc
	Args    []BulkStringNode
}

func (CommandNode) node() {}

func commandBuilder(data []byte) (Node, int, error) {
	sz, offset, err := extractArraySize(data)
	logger.Log.Info("Parsing array")
	if err != nil {
		logger.Log.Error("Failed to extract array size")
		return nil, 0, err
	}

	logger.Log.Info(fmt.Sprintf("array size is %d", sz))

	result := make([]BulkStringNode, sz)

	for i := range sz {
		internalData, delta, err := internalParse(data[offset:])

		if err != nil {
			logger.Log.Error(fmt.Sprintf("failed to extract data %d - %s", i, err))
			break
		}

		result[i] = internalData.(BulkStringNode)
		offset += delta
	}

	return CommandNode{Command: []BulkStringNode{result[0]}, Args: result[1:]}, 0, nil
}
