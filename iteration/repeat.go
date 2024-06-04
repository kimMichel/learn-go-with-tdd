package iteration

func Repeat(char string) string {
	var reps string
	for i := 0; i < 5; i++ {
		reps = reps + char
	}

	return reps
}