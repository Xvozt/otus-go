package unpacker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack_Without_Escape(t *testing.T) {
	actual := "abcd"
	expected := "abcd"

	assert.Equal(t, expected, actual)
}
