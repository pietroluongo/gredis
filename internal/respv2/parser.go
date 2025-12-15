package respv2

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/internal/logger"
)

func ParseMessage(rawData []byte) Node {
	stuff, _, err := internalParse(rawData)

	if err != nil {
		logger.Log.Error(fmt.Sprintf("Failed to parse message - %s", err.Error()))
		return SimpleStringNode{Data: "OK"} // TODO this should NOT be ok!
	}

	return stuff
}

func internalParse(rawData []byte) (Node, int, error) {
	trimmedMessage := cleanInput(rawData)

	if len(trimmedMessage) == 0 {
		logger.Log.Warn("message length is 0")
		return nil, 0, fmt.Errorf("message length is 0")
	}

	headerChar := trimmedMessage[0]

	builderFunction := handler_mapping[RespHeader(headerChar)]

	if builderFunction == nil {
		return nil, 0, fmt.Errorf("failed to match header %s on message %s trimmed to %s", string(headerChar), string(rawData), string(trimmedMessage))
	}

	return handler_mapping[RespHeader(headerChar)](trimmedMessage[1:])
}
