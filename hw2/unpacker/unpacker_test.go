package unpacker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack_Without_Escape_Two(t *testing.T) {
	actual, _ := Unpack("a4bc2d5e")
	expected := "aaaabccddddde"

	assert.Equal(t, expected, actual)
}

func TestUnpack_Empty_String(t *testing.T) {
	actual, err := Unpack("")
	expectedErrorMsg := "empty string"
	expected := ""

	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
	assert.Equal(t, actual, expected)
}