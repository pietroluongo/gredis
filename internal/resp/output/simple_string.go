package output

import "fmt"

func BuildSimpleString(s string) string {
	if len(s) == 0 {
		return "-1\r\n"
	}
	return fmt.Sprintf("+%s\r\n", s)
}
