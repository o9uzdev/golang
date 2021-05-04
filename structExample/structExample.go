package main

import "fmt"

type Post struct {
	title, body string
	index	string
}

func main() {
	post := Post{"Go Struct", "Go Struct Example", "1"}
	fmt.Println(post)

	post = Post{body: "Go Struct Example", index: "1", title: "Go Struct"}

	post = Post{}

	post = Post{title: "Go Struct Example"}

	var post2 Post
	post2.title = "Go Arrays Example"
	post2.index = "2"
	fmt.Println(post2)
}