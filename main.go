package main

import (
	"github.com/b1llywitant0/stamps-tech-test.git/generator"
)

func main() {
	const startNumber = 1
	const endNumber = 100
	generator.NewNumbersArray(startNumber, endNumber)
}
