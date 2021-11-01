/*
	Anagram adalah istilah dimana suatu string jika dibolak balik ordernya maka akan sama eg. 'aku' dan
	'kua' adalah Anagram, 'aku' dan 'aka' bukan Anagram.
	Dibawah ini ada array berisi sederetan Strings.
	['kita', 'atik', 'tika', 'aku', 'kia', 'makan', 'kua']
	Silahkan kelompokkan/group kata-kata di dalamnya sesuai dengan kelompok Anagramnya,
	# Expected Outputs
	[

	["kita", "atik", "tika"],
	["aku", "kua"],
	["makan"],
	["kia"]
	]
*/

package case4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Anagram(strs []string) [][]string {
	// alphabet lenght is 26 so the array length is 26
	m := map[[26]int][]string{}
	result := [][]string{}

	for _, val := range strs {
		var arr [26]int
		for i := range val {
			// the byte character of alphabet must be substract by byte of 'a'
			arr[val[i]-'a'] += 1
		}
		m[arr] = append(m[arr], val)
	}
	for i := range m {
		result = append(result, m[i])
	}

	return result
}

func TestAnagram(t *testing.T) {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	expectedResult := [][]string{{"kita", "atik", "tika"}, {"aku", "kua"}, {"kia"}, {"makan"}}
	result := Anagram(input)
	assert.ElementsMatch(t, expectedResult, result)
}
