package output

import "fmt"

func BuildSimpleError(s string) string {
	return fmt.Sprintf("-%s\r\n", s)
}
