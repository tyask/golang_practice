package main

import (
	"reflect"
	"strings"
	"testing"
)

/*
[演習]
BookSearchEngineのSearchWithPrefix/GroupByPublisherメソッドを実装してください。
*/

type Book struct {
	Title     string
	Outher    string
	Publisher string
	Price     int
}

type BookSearchEngine struct {
	books []Book
}

func (b *BookSearchEngine) SearchWithPrefix(prefix string) []Book {
	ret := []Book{}
	for _, book := range b.books {
		if strings.HasPrefix(book.Title, prefix) {
			ret = append(ret, book)
		}
	}
	return ret
}

func (b *BookSearchEngine) GroupByPublisher() map[string][]Book {
	ret := make(map[string][]Book)
	for _, book := range b.books {
		ret[book.Publisher] = append(ret[book.Publisher], book)
	}
	return ret
}

func TestSearchWithPrefix(t *testing.T) {
	b1 := Book{Title: "The Go Programming Language"}
	b2 := Book{Title: "Go Web Programming"}
	b3 := Book{Title: "Learning Go"}
	b4 := Book{Title: "Go Cookbook"}
	engine := BookSearchEngine{books: []Book{b1, b2, b3, b4}}

	actual := engine.SearchWithPrefix("Go")
	expected := []Book{b2, b4}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Should be %v, but %v", expected, actual)
	}
}

func TestGroupByPublisher(t *testing.T) {
	b1 := Book{Title: "The Go Programming Language", Publisher: "Addison-Wesley Professional"}
	b2 := Book{Title: "Go Web Programming", Publisher: "Manning Publications"}
	b3 := Book{Title: "Learning Go", Publisher: "O'Reilly Media, Inc."}
	b4 := Book{Title: "Go Cookbook", Publisher: "O'Reilly Media, Inc."}
	engine := BookSearchEngine{books: []Book{b1, b2, b3, b4}}

	actual := engine.GroupByPublisher()
	expected := map[string][]Book{
		"Addison-Wesley Professional": {b1},
		"Manning Publications":        {b2},
		"O'Reilly Media, Inc.":        {b3, b4},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Should be %v, but %v", expected, actual)
	}
}
