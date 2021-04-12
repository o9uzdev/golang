package main

import (
	"fmt"
	"os"
)

func getEnvironments() []string {
	return os.Environ()
}

func main() {
	for _, env := range getEnvironments() {
		fmt.Println(env)
	}
	
}
