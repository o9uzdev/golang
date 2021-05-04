package main

import "testing"

func TestMain(t *testing.T){
	post := Post{"Go Struct", "Go Struct Example", "1"}

	got1 := post.title
	want1 := "Go Struct"

	got2 := post.body
	want2 := "Go Struct Example"
	
	got3 := post.index
	want3 := "1"

	assertStrings(t, got1, want1)
	assertStrings(t, got2, want2)
	assertStrings(t, got3, want3)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}