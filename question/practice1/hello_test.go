package main

import (
	"fmt"
	"testing"
)

func Hello(name string) string {
	// TODO "Hello, {name}"の形式の文字列を返すように実装してください。
	return fmt.Sprintf("Hello, %s", name)
}

func TestHello(t *testing.T) {
	if Hello("Gopher") != "Hello, Gopher" {
		t.Error("Should be 'Hello, Gopher'")
	}
}
