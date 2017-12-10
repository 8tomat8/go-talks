package main

import (
	"fmt"
	"io"
)

type xStruct struct{}

func (xStruct) String() string { return "str" }

func (xStruct) Read(p []byte) (n int, err error) { return 0, nil }

func main() {
	type i1 interface {
		String() string
		Read(p []byte) (n int, err error)
	}
	var i i1 = xStruct{"Hodor"}
	i = i1(xStruct{"Hodor"})

	iReader := io.Reader(i) // iReader now of type "io.Reader"
	// iStringer now of type "fmt.Stringer". Do not omit ok, or it will panic in runtime!
	iStringer, ok := iReader.(fmt.Stringer)
	if !ok {
		panic("WAT!?")
	}

	//x := fmt.Stringer(iReader) // <-- But this will not compile

	// Assertion of source type. 
	sourceXStruct, ok := iStringer.(xStruct)
	if !ok {
		panic("...")
	}
	var x xStruct = sourceXStruct // sourceXStruct is now of type xStruct
	_ = x
}
