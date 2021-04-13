// https://play.golang.org/p/36WxyYWZcoN
package main

import (
	"fmt"
)

func main() {

	v := []string{"a", "b"}
	w := []string{"a", "b"}
	x := []string{"a", "b", "c"}

	y := CompareStringArray(v, w)
	z := CompareStringArray(w, x)

	fmt.Println(y)
	fmt.Println(z)
}

func CompareStringArray(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
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
