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
var ErrNotFunc = fmt.Errorf("received value is not a function")
var ErrMismatchedOutParams = fmt.Errorf("out params for the function does not match")

func ItReturn(fnPt interface{}, outs ...any) (e error) {
	reflectedPointerValue := reflect.ValueOf(fnPt)
	typeofFnPointer := reflectedPointerValue.Type()
	if typeofFnPointer.Kind() != reflect.Pointer {
		e = fmt.Errorf("%w: %v", ErrNotFunc, fnPt)
		return
	}
	ptAddy := fmt.Sprintf("%p", fnPt)

	fn := reflectedPointerValue.Elem()
	asInterface := fn.Interface()

	var mockedFn *MockedFunction

	// search for original value
	if storage, wasPreviouslyMocked := MockedHistory[ptAddy]; wasPreviouslyMocked {
		asInterface = storage.originalElement
		mockedFn = storage
	} else {
		mockedFn = &MockedFunction{
			fnPt:            fnPt,
			originalElement: asInterface,
		}
		MockedHistory[ptAddy] = mockedFn
	}

	typeofFn := fn.Type()
	if typeofFn.Kind() != reflect.Func {
		e = ErrNotFunc
		return
	}

	argTypes := make([]reflect.Type, typeofFn.NumIn())
	for i := 0; i < typeofFn.NumIn(); i++ {
		inArg := typeofFn.In(i)
		argTypes[i] = inArg
	}

	outTypes := make([]reflect.Type, typeofFn.NumOut())
	if len(outs) != typeofFn.NumOut() {
		e = ErrMismatchedOutParams
	}
	for i := 0; i < typeofFn.NumOut(); i++ {
		outTypes[i] = typeofFn.Out(i)
	}

	newReflectedFnType := reflect.FuncOf(argTypes, outTypes, typeofFn.IsVariadic())
	newReflectedFn := reflect.MakeFunc(newReflectedFnType, func(args []reflect.Value) (results []reflect.Value) {
		callArgs := make([]interface{}, len(args))
		for i, arg := range args {
			callArgs[i] = arg.Interface()
		}
		mockedFn.WasCalledWith(callArgs)

		results = make([]reflect.Value, len(outs))
		for i := 0; i < len(outs); i++ {
			results[i] = reflect.ValueOf(outs[i])
		}

		return
	})

	reflectedPointerValue.Elem().Set(newReflectedFn)
	return
}

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
