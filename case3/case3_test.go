/*
Please refactor the code below to make it more concise, efficient and readable with good logic flow.
func findFirstStringInBracket(str string) string {
	if (len(str) > 0) {
		indexFirstBracketFound := strings.Index(str,"(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket,")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound-1])
			}else{
				return ""
			}
			}else{
				eturn ""
			}
		}else{
			return ""
	}
	return ""
}
*/

package case3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	result := findFirstStringInBracket("bracket(test))")
	assert.Equal(t, "test", result)
}

func findFirstStringInBracket(str string) string {
	var result string
	if len(str) > 0 {
		indexFirstBracket := strings.Index(str, "(")
		if indexFirstBracket >= 0 {
			wordsAfterFirstBracket := str[indexFirstBracket:]
			indexClosingBracket := strings.Index(wordsAfterFirstBracket, ")") + indexFirstBracket
			if indexClosingBracket >= 0 {
				result = str[indexFirstBracket+1 : indexClosingBracket]
			}
		}
	}
	return result
}
