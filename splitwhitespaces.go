package piscine

func SplitWhiteSpaces(str string) []string {
	strC := []rune(str)

	length := 0
	for range str {
		length++
	}

	i := 0
	wordnumber := 1
	for i < length {
		if strC[i] == ' ' || strC[i] == '\t' || strC[i] == '\n' {
			for j := i + 1; j < length; j++ {
				if strC[j] == ' ' || strC[j] == '\t' || strC[j] == '\n' {
					continue
				} else {
					i = j - 1
					break
				}
			}
			wordnumber++
		}
		i++
	}

	result := make([]string, wordnumber)
	broken := false
	i = 0
	for index := range result {
		for i < length {
			if strC[i] == ' ' || strC[i] == '\t' || strC[i] == '\n' {
				for j := i + 1; j < length; j++ {
					if strC[j] == ' ' || strC[j] == '\t' || strC[j] == '\n' {
						continue
					} else {
						broken = true
						i = j
						break
					}
				}
			}
			if broken {
				broken = false
				break
			}
			result[index] += string(strC[i])
			i++
		}
	}
	return result
}
