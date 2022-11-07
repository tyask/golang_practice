package main

import (
	"reflect"
	"testing"
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

func (b *BookSearchEngine) SearchWithPrefixMulti(prefixes ...string) []Book {
	// TODO タイトルがprefixesのいずれかにマッチする本を検索し結果を返すように実装してください。
	// ある本が複数のprefixに引っかかったとしても、同じタイトルの本はただ1つだけ返すようにしてください。
	// また、探索結果は本のタイトルの昇順にソートしてください。
	return nil
}

func (b *BookSearchEngine) SearchWithPrefixMultiAsync(prefixes ...string) []Book {
	// TODO SearchWithPrefixMultiのマルチスレッド版を実装してください。
	// prefixesの各要素による探索処理をそれぞれ別スレッド(goroutine)で実行するようにし、最後にその結果を集約して返すようにしてください。
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
