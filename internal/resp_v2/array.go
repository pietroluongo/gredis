package respv2

import (
	"fmt"
)

func arrayBuilder(data []byte) (Node, int, error) {
	sz, offset, err := extractArraySize(data)
	log.Info("Parsing array")
	if err != nil {
		log.Error("Failed to extract array size")
		return nil, 0, err
	}

	log.Info(fmt.Sprintf("array size is %d", sz))

	result := make([]BulkStringNode, sz)

	for i := range sz {
		internalData, delta, err := internalParse(data[offset:])

		if err != nil {
			log.Error(fmt.Sprintf("failed to extract data %d - %s", i, err))
			break
		}

		result[i] = internalData.(BulkStringNode)
		offset += delta
	}

	return CommandNode{Command: result[0], Args: result[1:]}, 0, nil
}
