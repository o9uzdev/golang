package main

import (
	"fmt"
)

func Sum(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}

func main(){
	numbers := []int{3, 7, 3, 11, 22}
	fmt.Println(Sum(numbers))
}