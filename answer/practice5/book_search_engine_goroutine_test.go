package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"golang.org/x/exp/maps"
)

/*
[演習]
BookSearchEngineのSearchWithPrefixMulti/SearchWithPrefixMultiAsyncメソッドを実装してください。
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

func (b *BookSearchEngine) SearchWithPrefixMulti(prefixes ...string) []Book {
	m := map[string]Book{}
	for _, prefix := range prefixes {
		for _, book := range b.SearchWithPrefix(prefix) {
			m[book.Title] = book
		}
	}
	ret := maps.Values(m)
	sort.Slice(ret, func(i, j int) bool { return ret[i].Title < ret[j].Title })

	return ret
}

func (b *BookSearchEngine) SearchWithPrefixMultiAsync(prefixes ...string) []Book {
	ch := make(chan []Book)
	for _, prefix := range prefixes {
		go func(p string) {
			ch <- b.SearchWithPrefix(p)
		}(prefix)
	}

	m := map[string]Book{}
	for i := 0; i < len(prefixes); i++ {
		for _, book := range <-ch {
			m[book.Title] = book
		}
	}
	ret := maps.Values(m)
	sort.Slice(ret, func(i, j int) bool { return ret[i].Title < ret[j].Title })

	return ret
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
