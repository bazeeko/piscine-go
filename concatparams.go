package piscine

func ConcatParams(args []string) string {
	result := ""
	for i, word := range args {
		if i != 0 {
			result += "\n"
		}
		result += word
	}
	return result
}
