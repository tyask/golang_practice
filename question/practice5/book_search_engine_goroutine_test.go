package main

import (
	"reflect"
	"testing"
)

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
	return nil
}

func (b *BookSearchEngine) SearchWithPrefixMulti(prefixes ...string) []Book {
	return nil
}

func (b *BookSearchEngine) SearchWithPrefixMultiAsync(prefixes ...string) []Book {
	return nil
}

func TestSearchWithPrefixMulti(t *testing.T) {
	b1 := Book{Title: "The Go Programming Language"}
	b2 := Book{Title: "Go Web Programming"}
	b3 := Book{Title: "Learning Go"}
	b4 := Book{Title: "Go Cookbook"}
	engine := BookSearchEngine{books: []Book{b1, b2, b3, b4}}

	expected := []Book{b4, b2, b1}

	actual := engine.SearchWithPrefixMulti("The", "Go")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Should be %v, but %v", expected, actual)
	}

	actual2 := engine.SearchWithPrefixMultiAsync("The", "Go")
	if !reflect.DeepEqual(actual2, expected) {
		t.Errorf("Should be %v, but %v", expected, actual2)
	}
}
