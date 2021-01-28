package unpacker

import "fmt"

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("empty string")
	}
	return s, nil
}
