package respv2

type RespHeader byte

const (
	SimpleString RespHeader = '+'
	Command      RespHeader = '*'
	BulkString   RespHeader = '$'
)

type ParserFuncType func(data []byte) (Node, int, error)

var handler_mapping map[RespHeader]ParserFuncType

func init() {
	handler_mapping = map[RespHeader]ParserFuncType{
		SimpleString: simpleStringBuilder,
		Command:      commandBuilder,
		BulkString:   bulkStringBuilder,
	}
}
