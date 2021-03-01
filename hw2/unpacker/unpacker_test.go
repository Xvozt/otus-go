package unpacker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var invalidStringError = fmt.Errorf("invalid string")

type testData struct {
	input    string
	expected string
	err      error
}

func TestUnpack_for_test_data(t *testing.T) {
	testData := []testData{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		}, {
			input:    "abcd",
			expected: "abcd",
		}, {
			input:    "",
			expected: "",
		}, {
			input:    "45",
			expected: "",
			err:      invalidStringError,
		}, {
			input:    "aaa10b",
			expected: "",
			err:      invalidStringError,
		},
	}
	for _, data := range testData {
		result, err := Unpack(data.input)
		assert.Equal(t, data.expected, result)
		assert.Equal(t, data.err, err)
	}

}
