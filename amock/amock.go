package amock

import (
	"fmt"
	"reflect"
)

type MockedFunction struct {
	fnPt            interface{}
	originalElement interface{}
	callCount       int
	argsHistory     [][]interface{}
}

func (mf *MockedFunction) WasCalledWith(args []interface{}) {
	mf.callCount += 1
	mf.argsHistory = append(mf.argsHistory, args)
}

var MockedHistory = make(map[string]*MockedFunction)

func CallCount(fnPt interface{}) (count int, e error) {
	reflectedPointerValue := reflect.ValueOf(fnPt)
	typeofFnPointer := reflectedPointerValue.Type()
	if typeofFnPointer.Kind() != reflect.Pointer {
		e = fmt.Errorf("%w: %v", ErrNotFunc, fnPt)
		return
	}

	ptAddy := fmt.Sprintf("%p", fnPt)

	mockedFn, wasMocked := MockedHistory[ptAddy]
	if !wasMocked {
		return
	}

	count = mockedFn.callCount
	return
}
