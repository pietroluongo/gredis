package respv2

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/operations"
)

type Node interface {
	node()
}

type SimpleStringNode struct {
	Data string
}

func (SimpleStringNode) node() {}

type BulkStringNode struct {
	Data string
}

func (BulkStringNode) node() {}

type ArrayNode struct {
	Data []Node
}

func (ArrayNode) node() {}

type CommandNode struct {
	Command BulkStringNode
	Args    []BulkStringNode
}

func (CommandNode) node() {}

func (an ArrayNode) ToCommand() CommandNode {
	itemSize := len(an.Data)
	args := make([]BulkStringNode, itemSize-1)
	for i := range itemSize {
		args[i] = an.Data[i+1].(BulkStringNode)
	}
	return CommandNode{Command: an.Data[0].(BulkStringNode), Args: args}
}

func (n BulkStringNode) IsCommand() bool {
	return slices.Contains([]operations.ValidHandlers{operations.Ping, operations.Echo, operations.Get, operations.Set}, operations.ValidHandlers(strings.ToLower(n.Data)))
}

func (n SimpleStringNode) IsSimpleCommand() bool {
	return slices.Contains([]operations.ValidHandlers{operations.Ping}, operations.ValidHandlers(strings.ToLower(n.Data)))
}
