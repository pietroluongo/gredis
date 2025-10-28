package resp

type ValidHeaders byte

const (
	SimpleString ValidHeaders = '+'
	Array        ValidHeaders = '*'
	BulkString   ValidHeaders = '$'
)

var parseBuilders = map[ValidHeaders]func(data []byte) *RespMessage{
	SimpleString: simpleStringBuilder,
	Array:        arrayBuilder,
	BulkString:   todoBuilder,
}

func todoBuilder(data []byte) *RespMessage {
	log.Info("TODO HANDLER NOT IMPLEMENTED")
	return &RespMessage{Kind: SimpleString, Content: "OK"}
}
