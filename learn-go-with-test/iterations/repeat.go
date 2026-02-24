package iterations

import "strings"

func Repeat(ch string, c int) string {
	var repeated strings.Builder

	for range c {
		repeated.WriteString(ch)
	}

	return repeated.String()
}
