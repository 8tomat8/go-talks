package mock_example

import (
	"net/http/httptest"
	"testing"

	"encoding/json"
	"reflect"

	"bytes"

	"net/http"

	"github.com/ssOlexBaiko/library/api/web"
	"github.com/ssOlexBaiko/library/storage"
)

func TestGetBooks(t *testing.T) {
	store := &mockStorage{
		getBooks: storage.Books{storage.Book{"asd", "asd", []string{"1", "2"}, 123, 12.12}},
	}

	router := web.NewRouter(
		web.NewHandler(store),
	)

	req := httptest.NewRequest("GET", "/books", nil)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	resultBooks := storage.Books{}
	err := json.NewDecoder(rr.Body).Decode(&resultBooks)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resultBooks, store.getBooks) {
		t.Error("AAAA")
	}
}

func TestCreateBook(t *testing.T) {
	store := &mockStorage{
		createBook: storage.Books{storage.Book{ID: "Fus Roh dah!1", Title: "Not me =(", Genres: []string{"2", "1"}, Pages: 321, Price: 1111.11}},
	}
	router := web.NewRouter(
		web.NewHandler(store),
	)

	testBook := storage.Book{ID: "Fus Roh dah!", Title: "aHDFThaesfsd", Genres: []string{"1", "2"}, Pages: 123, Price: 12.12}

	buf := bytes.NewBuffer(make([]byte, 0, 2<<20))
	err := json.NewEncoder(buf).Encode(testBook)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/books", buf)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatal("smth does wrong")
	}

	if len(store.createBook) != 2 {
		t.Error("Book was lost")
	}

	if reflect.DeepEqual(store.createBook[0], store.createBook[1]) {
		t.Error("Nope")
	}

	if store.createBook[1].Pages != testBook.Pages {
		t.Error("Not equal")
	}

	if store.createBook[1].Title != testBook.Title {
		t.Error("Not equal")
	}

	if store.createBook[1].Price != testBook.Price {
		t.Error("Not equal")
	}

	if !reflect.DeepEqual(store.createBook[1].Genres, testBook.Genres) {
		t.Error("Not equal")
	}
}

type mockStorage struct {
	getBooks   storage.Books
	createBook storage.Books
}

func (m *mockStorage) GetBooks() (storage.Books, error) {
	return m.getBooks, nil
}

func (m *mockStorage) CreateBook(book storage.Book) error {
	book.PrepareToCreate()
	m.createBook = append(m.createBook, book)
	return nil
}

func (*mockStorage) GetBook(id string) (storage.Book, error) {
	panic("implement me")
}

func (*mockStorage) RemoveBook(id string) error {
	panic("implement me")
}

func (*mockStorage) ChangeBook(changedBook storage.Book) (storage.Book, error) {
	panic("implement me")
}

func (*mockStorage) PriceFilter(filter storage.BookFilter) (storage.Books, error) {
	panic("implement me")
}

func (*mockStorage) Close() error {
	return nil
}
