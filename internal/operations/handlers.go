package operations

type ValidHandlers string

const (
	Ping ValidHandlers = "ping"
	Echo ValidHandlers = "echo"
	Set  ValidHandlers = "set"
	Get  ValidHandlers = "get"
)
