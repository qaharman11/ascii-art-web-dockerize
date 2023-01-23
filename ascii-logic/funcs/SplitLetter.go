package funcs

func SplitLetter(str []string) []string {
	var letter []string
	c := 0
	temp := ""
	for i := 0; i < len(str); i++ {
		if str[i] != "\n" {
			temp += string(str[i]) + "\n"
			if c == 8 {
				letter = append(letter, temp)
				temp = ""
				c = 0
				i++
			}
			c++
		} else {
			continue
		}
	}
	letter[0] = letter[0][1:]
	return letter
}
