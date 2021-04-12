package main

import (
	"os"
	"testing"
)

func TestgetEnvironments(t *testing.T) {
	
	expected := os.Environ()
	actual := getEnvironments()
	
	if compareStringArray(expected, actual) {
		t.Error("Do not match")
	}

}

func compareStringArray(a, b []string) bool {

    // If one is nil, the other must also be nil.
    if (a == nil) != (b == nil) { 
        return false; 
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}