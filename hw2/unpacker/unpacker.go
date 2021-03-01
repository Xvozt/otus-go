package unpacker

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("empty string")
	}

	r := []rune(s)
	var result = strings.Builder{}

	for i := 0; i < len(r); i++ {
		if n, err := strconv.Atoi(string(r[i])); err == nil {
			if i == 0 || (i > 0 && unicode.IsDigit(r[i-1])) {
				return "", fmt.Errorf("invalid string")
			}
			result.WriteString(strings.Repeat(string(r[i-1]), n-1))
		} else {
			result.WriteRune(r[i])
		}
	}

	return result.String(), nil
}