package iteration

func Repeat(char string, times int) string {
	var reps string
	for i := 0; i < times; i++ {
		reps += char
	}

	return reps
}