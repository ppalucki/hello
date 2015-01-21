package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	got := simple()
	if got != "hello world" {
		t.Fail()
	}
}
