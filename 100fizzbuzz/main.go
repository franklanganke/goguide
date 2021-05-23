package main

import "fmt"

func main() {
	fmt.Println("main")

	for i := 1; i <= 100; i++ {
		fmt.Println("fizzBuzz(", i, ") == ", fb(i))
	}
}

func fb(i int) string {

	if isFizz(i) && isBuzz(i) {
		return "fizzBuzz"
	}

	if isFizz(i) {
		return "fizz"
	}

	if isBuzz(i) {
		return "buzz"
	}

	return fmt.Sprint(i)
}

func isFizz(i int) bool {
	if i%3 == 0 {
		return true
	}
	return false
}

func isBuzz(i int) bool {
	if i%5 == 0 {
		return true
	}
	return false
}
