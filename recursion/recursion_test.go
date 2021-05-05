package main

import "testing"

func TestFaktorial(t *testing.T){

	got := factorial(4)
	want := 24

	assertNumber(t, got, want)
}

func assertNumber(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}