package main

import (
	"testing"
)

func TestGetsString(t *testing.T) {
	want := "Hello world!"
	if got := getString(); got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}
