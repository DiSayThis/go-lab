package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseUTF8(t *testing.T) {
	assert.Equal(t, "olleh", ReverseUTF8("hello"))
	assert.Equal(t, "тевирП", ReverseUTF8("Привет"))
	assert.Equal(t, "!dlrow olleH", ReverseUTF8("Hello world!"))
}
