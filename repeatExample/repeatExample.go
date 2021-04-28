package main

import "fmt"

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 3; i++ {
		repeated = repeated + character
	}
	return repeated
}

func main(){
	fmt.Println(Repeat("a"))
}