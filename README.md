# amock

`amock` is a Golang package that provides assistance in monkey patching variables of type function.

# 1. Installation

To install `amock`, use the following `go get` command:

```
go get github.com/ye-yu/go-amock/amock
```

# 2. Usage

To use `amock`, you need to import the package in your Go code:

```go
package project

import "github.com/ye-yu/go-amock/amock"
```

Once you have imported `amock`, you can use its functions to monkey patch variables of type function.

The package provides three functions:

## 2.1. `amock.ItReturns`

```go
amock.ItReturns(funcPtr *interface{}, returnValues)
```

This function must take a pointer to a function that returns
value as the first parameter. The second parameter is a variadic
parameter for the values that should be returned by the function.


There is also a generic function implementation so that it
helps with the code completion and type checking.

```go
var GiveIntGetString = func(i int) s string {
	s = fmt.Sprintf("%d", i)
	return
}

println(GiveIntGetString(1)) // Outputs: "1"

amock.ItReturns1x1(&GiveIntGetString, "something else")

println(GiveIntGetString(1)) // Outputs: "something else"

```

## 2.2 `amock.CallCount`

This function returns the number of time the function is called.


# 3. Example

```go
package main_test

import (
	"fmt"
	"github.com/ye-yu/amock"
	amock2 "github.com/ye-yu/go-amock/amock"
	"testing"
)

func TestMockReturn(t *testing.T) {
	var GetInt = func() int {
		return 1
	}

	amock.ItReturns0x1(&GetInt, 2)

	println(GetInt())           // Output: 2
	println(amock2.CallCount(&GetInt)) // Output: 1
}
```