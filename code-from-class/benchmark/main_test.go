package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/RomanDerkach/homework/storage"
	"github.com/pkg/profile"
)

func BenchmarkJsonStore_Create(b *testing.B) {
	defer profile.Start().Stop()
	file, err := ioutil.TempFile(os.TempDir(), "jsonStorage")
	if err != nil {
		b.Fatal(err)
	}

	_, err = file.Write([]byte("[]"))
	if err != nil {
		b.Fatal(err)
	}

	fStat, err := file.Stat()
	if err != nil {
		b.Fatal(err)
	}

	store, err := storage.NewJSONStorage(os.TempDir() + "/" + fStat.Name())
	if err != nil {
		b.Fatal(err)
	}

	book := storage.Book{
		Price:  123.4576,
		Genres: []string{"a", "b", "c"},
		Pages:  1111,
		Title:  "Foooooooooooooooooooo",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := store.SaveBook(book)
		if err != nil {
			b.Error(err)
		}
	}
}

//
//func BenchmarkSQLStore_Create(b *testing.B) {
//
//	file, err := ioutil.TempFile(os.TempDir(), "sqlStorage")
//	if err != nil {
//		b.Fatal(err)
//	}
//
//	fStat, err := file.Stat()
//	if err != nil {
//		b.Fatal(err)
//	}
//
//	store, err := storage.NewSQLStorage(os.TempDir() + "/" + fStat.Name())
//	if err != nil {
//		b.Fatal(err)
//	}
//
//	book := storage.Book{
//		Price:  123.4576,
//		Genres: []string{"a", "b", "c"},
//		Pages:  1111,
//		Title:  "Foooooooooooooooooooo",
//	}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		err := store.SaveBook(book)
//		if err != nil {
//			b.Error(err)
//		}
//	}
//}
