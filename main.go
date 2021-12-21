package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("err: missing filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	filebyte, _ := os.ReadFile(filename)
	file := string(filebyte)

	tokens, err := tokenizeFile(file)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, t := range tokens {
		fmt.Printf("%+v\n", t)
	}
}
