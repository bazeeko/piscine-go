package piscine

func AtoiBase(s string, base string) int {
	str := []rune(s)
	baseC := []rune(base)
	baselength := RuneArrayLength(baseC)
	strlength := RuneArrayLength(str)
	for i := range baseC {
		if baseC[i] == '+' || baseC[i] == '-' {
			return 0
		}
		for j := i + 1; j < baselength; j++ {
			if baseC[i] == baseC[j] {
				return 0
			}
		}
	}
	if baselength < 2 {
		return 0
	}
	result := 0
	exp := 1
	for i := strlength - 1; i >= 0; i-- {
		index := 0
		for j := range baseC {
			if str[i] == baseC[j] {
				break
			}
			index++
		}
		result = result + index*exp
		exp = exp * baselength
	}
	return result
}
