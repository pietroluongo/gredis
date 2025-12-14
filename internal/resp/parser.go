package resp

import (
	"fmt"
)

type RespMessage struct {
	Kind    ValidHeaders
	Content any
}

func (r *RespMessage) IsArray() bool {
	return r.Kind == Array
}

func (r *RespMessage) AsMessageArray() ArrayRespMessage {
	return ArrayRespMessage{Kind: Array, Content: r.Content.([]RespMessage)}

}

func ParseMessage(rawData []byte) RespMessage {
	stuff, _, err := internalParse(rawData)

	if err != nil {
		log.Error(fmt.Sprintf("Failed to parse message - %s", err.Error()))
		return RespMessage{Kind: SimpleString, Content: "OK"}
	}

	return *stuff
}

func internalParse(rawData []byte) (*RespMessage, int, error) {
	trimmedMessage := cleanInput(rawData)

	if len(trimmedMessage) == 0 {
		log.Warn("message length is 0")
		return nil, 0, fmt.Errorf("message length is 0")
	}

	headerChar := trimmedMessage[0]
	builderFunction := parseBuilders[ValidHeaders(headerChar)]
	if builderFunction == nil {
		return nil, 0, fmt.Errorf("failed to match header %s on message %s trimmed to %s", string(headerChar), string(rawData), string(trimmedMessage))
	}
	return parseBuilders[ValidHeaders(headerChar)](trimmedMessage[1:])
}
