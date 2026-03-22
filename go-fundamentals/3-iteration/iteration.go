package iteration

func Repeat(character string) string {
	var repeated string

	// for range 5 {
	// 	repeated = repeated + character
	// }

	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}
