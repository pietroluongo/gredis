package output

import "fmt"

func BuildSimpleString(s string) string {
	return fmt.Sprintf("+%s\r\n", s)
}
