package generator

import "fmt"

func isPrime(number int) bool {
	if number <= 1 { // 1 and numbers below are not prime
		return false
	}
	if number <= 3 { // 2 and 3 are prime numbers (early-check)
		return true
	}
	if number%2 == 0 || number%3 == 0 { // Eliminate multiples of 2 and 3 (e.g., 4, 6, 8, 9, 10, 12, ...)
		return false
	}
	for i := 5; i*i <= number; i += 6 {
		// Only check numbers of the form 6n ± 1 (i and i+2)
		// These are the only possible prime candidates beyond 3
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}

func replaceFooBar(number int) interface{} {
	if number%3 == 0 && number%5 == 0 {
		return "FooBar"
	}
	if number%3 == 0 {
		return "Foo"
	}
	if number%5 == 0 {
		return "Bar"
	}
	return number
}

func NewNumbersArray(start int, end int) []interface{} {
	arr := []interface{}{}

	for i := start; i < end+1; i++ {
		if !isPrime(i) {
			if replaceFooBar(i) == "FooBar" {
				arr = append(arr, "FooBar")
				if len(arr) == 1 {
					fmt.Print("FooBar")
				} else {
					fmt.Print(" ", "FooBar")
				}
			} else if replaceFooBar(i) == "Bar" {
				arr = append(arr, "Bar")
				if len(arr) == 1 {
					fmt.Print("Bar")
				} else {
					fmt.Print(" ", "Bar")
				}
			} else if replaceFooBar(i) == "Foo" {
				arr = append(arr, "Foo")
				if len(arr) == 1 {
					fmt.Print("Foo")
				} else {
					fmt.Print(" ", "Foo")
				}
			} else {
				arr = append(arr, i)
				if len(arr) == 1 {
					fmt.Print(i)
				} else {
					fmt.Print(" ", i)
				}
			}
		}
	}

	return arr
}
