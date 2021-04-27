package main

import "fmt"

type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
	return d[word]
}

func main(){
	dictionary := Dictionary{"code": "this is just a code"}
	fmt.Println(dictionary.Search("code"))
}