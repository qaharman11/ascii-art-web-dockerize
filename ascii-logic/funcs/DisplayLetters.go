package funcs

func DisplayLetters(table []int, char []string, letters []string) string {
	var buf []int
	var str, s string
	for i := 0; i < len(table); i++ {
		if table[i] != 1000 {
			buf = append(buf, table[i])
		} else {
			if LetterOfIndex(buf, letters) == nil {
				s += "\n"
				continue
			} else {
				s += WriteText(LetterOfIndex(buf, letters))
				buf = nil
			}
			// s += "\n"
		}
	}
	if LetterOfIndex(buf, letters) != nil {
		LetterOfIndex(buf, letters)
		str = WriteText(LetterOfIndex(buf, letters))
	}
	ss := s + str
	return ss
}
