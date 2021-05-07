package main

import "fmt"

func main() {
	alphabet := map[string]string{
		"a": "alpha",
		"b": "bravo",
		"c": "charlie",
		"d": "delta",
		"e": "echo",
		"f": "foxtrott",
	}

	printAlphabet(alphabet)
}

func printAlphabet(a map[string]string) {
	for char, word := range a {
		fmt.Printf("%v --> %v\n", char, word)
	}
}
