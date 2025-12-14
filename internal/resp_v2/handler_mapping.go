package respv2

type ValidHeaders byte

const (
	SimpleString ValidHeaders = '+'
	Array        ValidHeaders = '*'
	BulkString   ValidHeaders = '$'
)

type ParserType func(data []byte) (Node, int, error)

var handler_mapping map[ValidHeaders]ParserType

func init() {
	handler_mapping = map[ValidHeaders]ParserType{
		SimpleString: simpleStringBuilder,
		Array:        arrayBuilder,
		BulkString:   bulkStringBuilder,
	}
}
