package funcs

func Logic(InputText string, banner string) (string, error) {
	filedata, err := Readfile(banner)
	if err != nil {
		return "", err
	}	
	letters := SplitLetter(filedata)
	Input, err := CheckAscii(InputText)
	if err != nil{
		return "", err
	}
	table := IndexesInTable(Input)
	char := LetterOfIndex(table, letters)
	output := DisplayLetters(table, char, letters)
	return output, nil
}
