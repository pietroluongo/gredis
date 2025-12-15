package respv2

type SimpleStringNode struct {
	Data string
}

func (SimpleStringNode) node() {}

func simpleStringBuilder(data []byte) (Node, int, error) {
	return SimpleStringNode{Data: string(data)}, len(data), nil
}
