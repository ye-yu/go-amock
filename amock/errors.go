package amock

import (
	"fmt"
)

var ErrNotFunc = fmt.Errorf("received value is not a function")
var ErrMismatchedOutParams = fmt.Errorf("out params for the function does not match")
