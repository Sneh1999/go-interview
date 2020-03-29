package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, Encode(2, 2), "10")
	assert.Equal(t, Encode(10, 10), "10")
	assert.Equal(t, Encode(10, 8), "12")
	assert.Equal(t, Encode(10, 16), "A")
	assert.Equal(t, Encode(12, 2), "1100")
	assert.Equal(t, Encode(282, 16), "11A")
	assert.Equal(t, Encode(83, 8), "123")
}

func TestDecode(t *testing.T) {
	assert.Equal(t, Decode("10", 2), 2)
	assert.Equal(t, Decode("10", 10), 10)
	assert.Equal(t, Decode("12", 8), 10)
	assert.Equal(t, Decode("A", 16), 10)
	assert.Equal(t, Decode("1100", 2), 12)
	assert.Equal(t, Decode("11A", 16), 282)
	assert.Equal(t, Decode("123", 8), 83)
}