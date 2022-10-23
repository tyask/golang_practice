package main

import (
	"fmt"
	"testing"
)

type Book struct {
	Title     string
	Outher    string
	Publisher string
	Price     int
}

func (b *Book) String() string {
	return fmt.Sprintf("%s/%s/%s (¥%d)", b.Title, b.Outher, b.Publisher, b.Price)
}

func (b *Book) SetPrice(price int) {
	b.Price = price
}

func TestStringForBook(t *testing.T) {
	b := Book{Title: "aaa", Outher: "bbb", Publisher: "ccc", Price: 1000}
	if b.String() != "aaa/bbb/ccc (¥1000)" {
		t.Error("Should be 'aaa/bbb/ccc (¥1000)'")
	}
}

func TestSetPriceForBook(t *testing.T) {
	b := Book{Price: 1000}
	b.SetPrice(2000)
	if b.Price != 2000 {
		t.Error("Should be 2000")
	}
}
