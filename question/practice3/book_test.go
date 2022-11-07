package main

import (
	"testing"
)

/*
[演習]
Book構造体にフィールドを定義し、String/SetPriceメソッドを実装してください。
*/

type Book struct {
	// TODO 以下のフィールドを定義してください。
	// Title     (string)
	// Outher    (string)
	// Publisher (string)
	// Price     (int)
}

func ( /* TODO receiverを定義してください */ ) String() string {
	// TODO 以下の形式でBook構造体の文字列表現を返すように実装してください。
	// {Title}/{Outher}/{Publisher} ({Price})
	return ""
}

func ( /* TODO receiverを定義してください */ ) SetPrice(price int) {
	// TODO Book.Priceに値を設定するように実装してください。
}

func TestStringForBook(t *testing.T) {
	b := Book{
		Title:     "The Go Programming Language",
		Outher:    "Alan Donovan, Brian Kernighan",
		Publisher: "Addison-Wesley Professional",
		Price:     2336,
	}
	expected := "The Go Programming Language/Alan Donovan, Brian Kernighan/Addison-Wesley Professional (2336)"
	if b.String() != expected {
		t.Errorf("Should be '%s'", expected)
	}
}

func TestSetPriceForBook(t *testing.T) {
	b := Book{Price: 1000}
	b.SetPrice(2000)
	if b.Price != 2000 {
		t.Error("Should be 2000")
	}
}
