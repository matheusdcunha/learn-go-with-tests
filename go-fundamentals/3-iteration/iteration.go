package iteration

import "strings"

func Repeat(character string, times int) string {
	var repeated strings.Builder

	// for range 5 {
	// 	repeated = repeated + character
	// }

	for i := 0; i < times; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
