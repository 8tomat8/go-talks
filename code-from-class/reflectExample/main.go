package main

import (
	"fmt"
	"reflect"
)

func main() {
	book := &Book{
		"IDVal",
		"TitleVal",
		[]string{"Ganres"},
		12,
		88.88,
	}

	bookRef := reflect.ValueOf(book)
	bookRef = reflect.Indirect(bookRef)

	switch bookRef.Kind() {
	case reflect.Slice:
		fmt.Println(bookRef.Len())
	case reflect.Struct:
		for i := 0; i < bookRef.NumField(); i++ {
			fmt.Println(bookRef.Field(i).Interface())
		}
	}

}

type Book struct {
	ID     string   `json:"id, omitempty"`
	Title  string   `json:"title, omitempty"`
	Ganres []string `json:"ganres, omitempty"`
	Pages  int      `json:"pages, omitempty"`
	Price  float64  `json:"price, omitempty"`
}
