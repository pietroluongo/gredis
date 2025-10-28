package resp

func simpleStringBuilder(data []byte) *RespMessage {
	return &RespMessage{Kind: ValidHeaders(SimpleString), Content: string(data)}
}
