package output

import "fmt"

func BuildBulkString(s string) string {
	if len(s) == 0 {
		return "$-1\r\n"
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
}
