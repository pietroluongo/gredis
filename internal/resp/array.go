package resp

import (
	"fmt"
)

type ArrayRespMessage struct {
	Kind    ValidHeaders
	Content []RespMessage
}

func arrayBuilder(data []byte) (*RespMessage, int, error) {
	sz, offset, err := extractArraySize(data)
	log.Info("Parsing array")
	if err != nil {
		log.Error("Failed to extract array size")
		return nil, 0, err
	}

	log.Info(fmt.Sprintf("array size is %d", sz))

	result := make([]RespMessage, sz)

	for i := range sz {
		data, delta, err := internalParse(data[offset:])
		if err != nil {
			log.Error(fmt.Sprintf("failed to extract data %d - %s", i, err))
			break
		}
		result[i] = *data
		offset += delta
	}
	return &RespMessage{Kind: Array, Content: result}, 0, nil
}
