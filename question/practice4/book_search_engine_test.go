package main

import (
	"reflect"
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
	// TODO タイトルの前方一致で検索するように実装してください。
	// 前方一致は、strings.HasPrefixメソッドを利用してください。
	return nil
}

func (b *BookSearchEngine) GroupByPublisher() map[string][]Book {
	// TODO 本の一覧をPublisherごとに分類して返すように実装してください。
	return nil
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
