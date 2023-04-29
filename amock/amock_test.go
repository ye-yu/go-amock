package amock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItReturn(t *testing.T) {
	var ToMockReturn = func() int {
		return 1
	}
	ItReturns(&ToMockReturn, 2)

	out := ToMockReturn()

	assert.Equal(t, 2, out)
}

func TestCallCount(t *testing.T) {
	var ToMockReturn = func() int {
		return 1
	}
	ItReturns(&ToMockReturn, 2)

	out := ToMockReturn()

	assert.Equal(t, 2, out)

	callCount, e := CallCount(&ToMockReturn)
	if e != nil {
		t.Fatal(e)
		return
	}
	assert.Equal(t, 1, callCount)
}
