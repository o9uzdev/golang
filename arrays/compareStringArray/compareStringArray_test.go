package main

import (
	"testing"
)

func TestCompareStringArray(t *testing.T) {

	CompareTwoNilArrayExpectTrue := CompareStringArray([]string{}, []string{})
	if !CompareTwoNilArrayExpectTrue {
		t.Error("Compared two nil array and returned false")
	}

	AtLeastOneNilExpectFalse := CompareStringArray([]string{}, []string{"a", "b"})
	if AtLeastOneNilExpectFalse {
		t.Error("If one is nil, the other must also be nil")
	}

	CompareTwoArrayExpectFalse := CompareStringArray([]string{"a", "b"}, []string{"a", "b", "c"})
	if CompareTwoArrayExpectFalse {
		t.Error("Compared different two array and returned true")
	}

	CompareTwoArrayExpectTrue := CompareStringArray([]string{"a", "b", "c"}, []string{"a", "b", "c"})
	if !CompareTwoArrayExpectTrue {
		t.Error("Compared different two array and returned false")
	}

}
