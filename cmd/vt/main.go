package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/iamleot/go-treccani"
	"github.com/mitchellh/go-wordwrap"
)

// PrintDefinition pretty prints a term definition.
func PrintDefinition(definition string) {
	fmt.Println(wordwrap.WrapString(definition, 80))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s term\n", os.Args[0])
		os.Exit(1)
	}

	term := os.Args[1]

	for i, definition := range treccani.Terms(term, http.DefaultClient) {
		if i > 0 {
			fmt.Printf("\n")
		}
		PrintDefinition(definition)
	}
}
