package main

import (
	"io"

	"time"

	"log"

	"fmt"

	"github.com/pborman/uuid"
)

type uds struct {
	a string
	b int
	c interface{}
}

func test(v interface{}) string {
	rwc, ok := v.(io.ReadWriteCloser)
	if !ok {
		return "Oops!"
	}
	_ = rwc
	return "Asserted!"
}

func main() {
	//testChan := make(chan string)
	//
	//fmt.Println("nil:", test(nil))
	//fmt.Println("string:", test("string"))
	//fmt.Println("123:", test(123))
	//fmt.Println("0.123:", test(0.123))
	//fmt.Println("chan string:", test(testChan))
	//fmt.Println("custom struct:", test(uds{}))
	//fmt.Println("[]int:", test([]int{1, 2, 3, 4}))
	//fmt.Println("io.ReadWriteCloser:", test(&os.File{}))

	ifImplement(&User{age: 12, name: "asd", legs: []int{20, 50}})
}

func ifImplement(v interface{}) {
	obj, ok := v.(NewIDer)
	if ok {
		fmt.Println("Woohoo!!!")
		if err := obj.NewID(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("%#v\n", v)

}

type NewIDer interface {
	NewID() error
}

type User struct {
	id       uuid.UUID
	name     string
	age      int
	legs     []int
	createAt time.Time
}

func (u User) Name() string {
	u.name = "AAAA"
	return u.name
}

func (u *User) Age() int {
	return u.age
}

func (u *User) NewID() error {
	u.id = uuid.NewUUID()
	return nil
}
