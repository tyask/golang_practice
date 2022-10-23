package main

import (
	"reflect"
	"strings"
	"testing"
)

type Book struct {
	Title     string
	Outher    string
	Publisher string
	Price     int
}

type Booble struct {
	books []Book
}

func (b *Booble) SearchWithPrefix(prefix string) []Book {
	ret := []Book{}
	for _, book := range b.books {
		if strings.HasPrefix(book.Title, prefix) {
			ret = append(ret, book)
		}
	}
	return ret
}

func (b *Booble) GroupByPublisher() map[string][]Book {
	ret := make(map[string][]Book)
	for _, book := range b.books {
		ret[book.Publisher] = append(ret[book.Publisher], book)
	}
	return ret
}

func TestSearchWithPrefix(t *testing.T) {
	b1 := Book{Title: "aa"}
	b2 := Book{Title: "ax"}
	b3 := Book{Title: "bb"}
	b4 := Book{Title: "bx"}
	booble := Booble{books: []Book{b1, b2, b3, b4}}

	actual := booble.SearchWithPrefix("a")
	expected := []Book{b1, b2}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Should be %v, but %v", expected, actual)
	}
}

func TestGroupByPublisher(t *testing.T) {
	b1 := Book{Title: "aa", Publisher: "abc"}
	b2 := Book{Title: "ax", Publisher: "xyz"}
	b3 := Book{Title: "bb", Publisher: "abc"}
	b4 := Book{Title: "bx", Publisher: "xyz"}
	booble := Booble{books: []Book{b1, b2, b3, b4}}

	actual := booble.GroupByPublisher()
	expected := map[string][]Book{
		"abc": {b1, b3},
		"xyz": {b2, b4}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Should be %v, but %v", expected, actual)
	}
}
