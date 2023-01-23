package funcs

import (
	"bufio"
	"fmt"
	"os"
)

func Readfile(filename string) ([]string, error) {
	var letters []string
	path := "ascii-logic/banners/"
	tfile := ".txt"
	final := path + filename + tfile
	hmap := map[string]string{
		"standard":   "ac85e83127e49ec42487f272d9b9db8b",
		"shadow":     "95004356a0153fc02470b664ade53b57",
		"thinkertoy": "bf1d925662e40f5278b26a0531bfdb63",
	}

	file, err := os.Open(final)
	if err != nil {
		return letters, err
	}
	if hmap[filename] != Md5sum(final) {
		return letters, fmt.Errorf("file has benn changed")
	}
	Scanner := bufio.NewScanner(file)

	for Scanner.Scan() {
		letters = append(letters, string(Scanner.Bytes()))
	}

	return letters, nil
}

func CheckAscii(letters string) (string, error) {
	for i := 0; i < len(letters); i++ {
		if letters[i] >= 0 && letters[i] <= 126 {
			continue
		}
		return "", fmt.Errorf("out of ascii characters")
	}
	return letters, nil
}
