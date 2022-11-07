package main

import (
	"fmt"
	"testing"
)

/*
[演習]
Book構造体にフィールドを定義し、String/SetPriceメソッドを実装してください。
*/

type Book struct {
	Title     string
	Outher    string
	Publisher string
	Price     int
}

func (b *Book) String() string {
	return fmt.Sprintf("%s/%s/%s (%d)", b.Title, b.Outher, b.Publisher, b.Price)
}

func (b *Book) SetPrice(price int) {
	b.Price = price
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
