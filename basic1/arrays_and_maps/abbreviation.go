package arrays_and_maps

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func MakeAbbreviation() {
	phrase := readString()
	words := strings.Fields(phrase)

	var abbr string
	for _, word := range words {
		char := []rune(word[0:1])

		if char[0] > 122 {
			char = []rune(word[0:2])
		}

		var letter rune
		if unicode.IsLetter(char[0]) {
			letter = unicode.ToUpper(char[0])
		} else {
			continue
		}

		abbr += string(letter)
	}

	fmt.Println(string(abbr))
}

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
