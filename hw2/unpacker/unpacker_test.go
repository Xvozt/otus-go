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
	for _, v := range []testData{
		{
			input: "a4bc2d5e",
			expected: "aaaabccddddde",
		},{
			input: "abcd",
			expected: "abcd",
		},{
			input: "",
			expected: "",
		},{
			input: "45",
			expected: "",
			err: invalidStringError,
		},{
			input: "aaa10b",
			expected: "",
			err: invalidStringError,
		},
	} {
		result, err := Unpack(v.input)
		assert.Equal(t, v.expected, result)
		assert.Equal(t, v.err, err)
	}



}