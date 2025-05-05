package main

import (
	"fmt"

	"github.com/b1llywitant0/stamps-tech-test.git/generator"
)

func main() {
	arr := generator.NewNumbersArray(1, 100)
	fmt.Println(arr)
}
