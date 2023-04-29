package amock

import (
	"fmt"
	"reflect"
)

type MockStorage struct {
	fnPt            interface{}
	originalElement interface{}
}

var MockedHistory = []MockStorage{}
var ErrNotFunc = fmt.Errorf("received value is not a function")
var ErrMismatchedOutParams = fmt.Errorf("out params for the function does not match")

func ItReturn(fnPt interface{}, outs ...any) (e error) {
	reflectedPointerValue := reflect.ValueOf(fnPt)
	typeofFnPointer := reflectedPointerValue.Type()
	if typeofFnPointer.Kind() != reflect.Pointer {
		e = fmt.Errorf("%w: %v", ErrNotFunc, fnPt)
		return
	}

	fn := reflectedPointerValue.Elem()
	asInterface := fn.Interface()

	// search for original value
	var wasPreviouslyMocked = false
	for _, storage := range MockedHistory {
		if storage.fnPt == fnPt {
			asInterface = storage.originalElement
			wasPreviouslyMocked = true
		}
	}

	// if it is not previously mocked, register in mock storage
	if !wasPreviouslyMocked {
		MockedHistory = append(MockedHistory, MockStorage{
			fnPt:            fnPt,
			originalElement: asInterface,
		})
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
		//currentCount := context.fnCountIndex.Get(int(spyToken))
		//context.fnCountIndex.Set(int(spyToken), currentCount+1)
		//
		//callHistoryStack := context.callingHistory.Get(int(spyToken))
		//if callHistoryStack == nil {
		//	callHistoryStack = [][]interface{}{}
		//}
		//argsStream := stream.OfSlice(args)
		//callArgs := stream.Map(
		//	argsStream,
		//	func(it reflect.Value) interface{} {
		//		return it.Interface()
		//	},
		//).ToSlice()
		//callHistoryStack = append(callHistoryStack, callArgs)

		results = make([]reflect.Value, len(outs))
		for i := 0; i < len(outs); i++ {
			results[i] = reflect.ValueOf(outs[i])
		}

		return
	})

	reflectedPointerValue.Elem().Set(newReflectedFn)
	return
}
