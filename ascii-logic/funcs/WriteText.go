package funcs

func WriteText(char []string) string {
	l := 8
	c := 0
	str := ""
	for l != 0 {
		for i := 0; i < len(char); i++ {
			row := len(char[i]) / 8
			for j := row * c; j < len(char[i]); j++ {
				if char[i][j] != '\n' {
					// fmt.Print(string(char[i][j]))
					str += string(char[i][j])
				} else {
					break
				}
			}
		}
		str += "\n"
		l--
		c++
	}
	return str
}
