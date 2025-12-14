package respv2

import "fmt"

func ParseMessage(rawData []byte) Node {
	stuff, _, err := internalParse(rawData)

	if err != nil {
		log.Error(fmt.Sprintf("Failed to parse message - %s", err.Error()))
		return SimpleStringNode{Data: "OK"}
	}

	return stuff
}

func internalParse(rawData []byte) (Node, int, error) {
	trimmedMessage := cleanInput(rawData)

	if len(trimmedMessage) == 0 {
		log.Warn("message length is 0")
		return nil, 0, fmt.Errorf("message length is 0")
	}

	headerChar := trimmedMessage[0]
	builderFunction := handler_mapping[ValidHeaders(headerChar)]
	if builderFunction == nil {
		return nil, 0, fmt.Errorf("failed to match header %s on message %s trimmed to %s", string(headerChar), string(rawData), string(trimmedMessage))
	}
	return handler_mapping[ValidHeaders(headerChar)](trimmedMessage[1:])
}
