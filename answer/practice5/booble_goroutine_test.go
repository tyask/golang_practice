package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"golang.org/x/exp/maps"
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

func (b *Booble) SearchWithPrefixMulti(prefixes ...string) []Book {
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

func (b *Booble) SearchWithPrefixMultiAsync(prefixes ...string) []Book {
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
	b1 := Book{Title: "aa"}
	b2 := Book{Title: "ax"}
	b3 := Book{Title: "bb"}
	b4 := Book{Title: "bx"}
	booble := Booble{books: []Book{b1, b2, b3, b4}}

	expected := []Book{b1, b2, b3}

	actual := booble.SearchWithPrefixMulti("a", "bb")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Should be %v, but %v", expected, actual)
	}

	actual2 := booble.SearchWithPrefixMultiAsync("a", "bb")
	if !reflect.DeepEqual(actual2, expected) {
		t.Errorf("Should be %v, but %v", expected, actual2)
	}
}
