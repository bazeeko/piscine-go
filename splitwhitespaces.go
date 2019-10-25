package piscine

func SplitWhiteSpaces(str string) []string {
	strC := []rune(str)

	length := 0
	for range str {
		length++
	}

	wordnumber := 1
	for _, char := range strC {
		if char == '\n' || char == '\t' || char == ' ' {
			wordnumber++
		}
	}

	result := make([]string, wordnumber)

	index := 0

	for i := range result {
		result[i] = ""
		for j := index; j < length; j++ {
			if strC[j] == '\n' || strC[j] == '\t' || strC[j] == ' ' {
				index = j + 1
				break
			}
			result[i] += string(strC[j])
		}
	}
	return result
}
