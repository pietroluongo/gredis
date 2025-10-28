package resp

import "fmt"

type RespMessage struct {
	Kind    ValidHeaders
	Content any
}

func ParseMessage(rawData []byte) RespMessage {
	trimmedMessage := cleanInput(rawData)

	if len(trimmedMessage) == 0 {
		log.Warn("message length is 0")
		return RespMessage{Kind: SimpleString, Content: "OK"}
	}

	stuff, err := internalParse(trimmedMessage)

	if err != nil {
		log.Error(fmt.Sprintf("Failed to parse message - %s", err.Error()))
		return RespMessage{Kind: SimpleString, Content: "OK"}
	}

	return *stuff
}

func internalParse(data []byte) (*RespMessage, error) {
	headerChar := data[0]
	builderFunction := parseBuilders[ValidHeaders(headerChar)]
	if builderFunction == nil {
		log.Error(fmt.Sprintf("Failed to match header %s", string(headerChar)))
		return nil, fmt.Errorf("failed to parse")
	}
	stuff := parseBuilders[ValidHeaders(headerChar)](data[1:])
	return stuff, nil
}
