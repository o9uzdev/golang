package main

import (
	"fmt"
	"strconv"
)

func map2json(from map[string]int) string {
	result := "{"
	mlen := len(from)
	i := 1
	for key, element := range from {
		t := strconv.Itoa(element)
		result = result + "\"" + key + "\"" + ":" + t
		if i != mlen {
			result = result + ","
		}
		i += 1

	}
	return result + "}"
}

func main() {

	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5

	fmt.Println("map:", map2json(m))
	fmt.Println("map:", m)
}