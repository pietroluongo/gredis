package respv2

func simpleStringBuilder(data []byte) (Node, int, error) {
	return SimpleStringNode{Data: string(data)}, len(data), nil
}
