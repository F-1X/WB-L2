package raspakouka

import (
	"fmt"
	"strconv"
	"strings"
)

func Raspokouka(str string) (string, error) {

	if str == "" {
		return "", nil
	}

	// output := make([]rune, 50)
	var output strings.Builder
	// len of output

	runs := []rune(str)


	for i := 0; i < len(runs); i++ {

		if !isDigit(runs[i]) && !isSlash(runs[i]) {
			output.WriteRune(runs[i])
			continue
		}

		if isSlash(runs[i]) {
		
			if i != len(runs)-1 {
				output.WriteRune(runs[i])
			} else {
				return "", fmt.Errorf("incorrect slash")
			}
			continue
		}

		if isDigit(runs[i]) {
			// if !escape {
			// 	return "", fmt.Errorf("incorrect digit %s after %s with no escape: %s, result: %s", string(runs[i]), string(runs[i-1]), str, output.String())
			// }
			count, _ := strconv.Atoi(string(runs[i]))
			output.WriteString(strings.Repeat(string(runs[i-1]), count))
	
		}

	}
	return output.String(), nil
}

var alphabet map[rune]struct{}

func init() {
	alphabet = map[rune]struct{}{'0': {}, '1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {}}
}

func isDigit(r rune) bool {
	if _, ok := alphabet[r]; ok {
		return true
	}
	return false
}

func isSlash(r rune) bool {
	return string(r) == "\\"
}
