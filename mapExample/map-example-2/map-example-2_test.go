package main

import "testing"

func Testmap2json(t *testing.T) {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5

	got := map2json(m)
	want := "{\"two\":2,\"three\":3,\"four\":4,\"five\":5,\"one\":1}"
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}