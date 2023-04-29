package amock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItReturns0x1(t *testing.T) {
	var GetInt = func() int {
		return 1
	}
	ItReturns0x1(&GetInt, 2)
	assert.Equal(t, 2, GetInt())
}

func TestItReturns9x9(t *testing.T) {
	var GetInt = func(
		string,
		string,
		string,
		string,
		string,
		string,
		string,
		string,
		string,
	) (
		int,
		int,
		int,
		int,
		int,
		int,
		int,
		int,
		int,
	) {
		return 1, 2, 3, 4, 5, 6, 7, 8, 9
	}
	ItReturns9x9(&GetInt, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	out1, out2, out3, out4, out5, out6, out7, out8, out9 := GetInt(
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	)
	assert.Equal(t, 11, out1)
	assert.Equal(t, 12, out2)
	assert.Equal(t, 13, out3)
	assert.Equal(t, 14, out4)
	assert.Equal(t, 15, out5)
	assert.Equal(t, 16, out6)
	assert.Equal(t, 17, out7)
	assert.Equal(t, 18, out8)
	assert.Equal(t, 19, out9)
}
