package resp

func simpleStringBuilder(data []byte) (*RespMessage, int, error) {
	return &RespMessage{Kind: ValidHeaders(SimpleString), Content: string(data)}, len(data), nil
}
