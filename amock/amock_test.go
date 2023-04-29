package amock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItReturn(t *testing.T) {
	var ToMockReturn = func() int {
		return 1
	}
	ItReturn(&ToMockReturn, 2)

	out := ToMockReturn()

	assert.Equal(t, 2, out)
}
