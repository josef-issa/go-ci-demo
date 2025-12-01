package main

import "testing"

func TestDummy(t *testing.T) {
	got := 1 + 1
	want := 2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
