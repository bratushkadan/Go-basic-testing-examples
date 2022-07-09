package util

import "fmt"

func Concat(strings ...fmt.Stringer) string {
	output := ""

	for _, stringer := range strings {
		output += stringer.String()
	}

	return output
}
