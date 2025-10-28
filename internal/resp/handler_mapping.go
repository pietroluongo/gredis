package resp

type ValidHeaders byte

const (
	SimpleString ValidHeaders = '+'
	Array        ValidHeaders = '*'
	BulkString   ValidHeaders = '$'
)

type ParserType func(data []byte) (*RespMessage, int, error)

var parseBuilders map[ValidHeaders]ParserType

func init() {
	parseBuilders = map[ValidHeaders]ParserType{
		SimpleString: simpleStringBuilder,
		Array:        arrayBuilder,
		BulkString:   bulkStringBuilder,
	}
}
