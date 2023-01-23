package funcs

func IndexesInTable(s string) []int {
	// array of indexes of input text
	var table []int
	var buf int
	str := []rune(s)
	// recording indexes to array
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			buf = 1000
			table = append(table, buf)
			buf = 0
			// i += 1
			continue
		} else {
			buf = int(str[i]) - 32
			table = append(table, buf)
			buf = 0
		}
	}

	return table
}

func LetterOfIndex(table []int, letters []string) []string {
	// array of string that contains each letter's ascii symbols
	var char []string

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(letters); j++ {
			if table[i] == j {
				char = append(char, letters[j])
			}
		}
	}

	return char
}
 