package main

import (
	"fmt"
	"log"
)

type ErrBar string

func (ErrBar) Error() string {
	return "ErrBar ERROR"
}

func main() {
	//var err *ErrBar = nil
	err := Foo()    //((nil), (*Type(*ErrBar)))
	if err != nil { //((nil), (nil))
		fmt.Println("Not nil")
		log.Fatal(err)
	}
	fmt.Println("Succsess")
}

//
//func Foo() error {
//	return errors.New("ERROR!!!")
//}

func Foo() error {
	var err *ErrBar = nil
	return err
}
