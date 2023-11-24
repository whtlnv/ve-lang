package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "no file path provided\n")
		return
	}

	filePath := args[0]

	if filePath == "" {
		fmt.Fprintf(os.Stderr, "no file path provided\n")
		return
	}

	fmt.Println("reading from:", filePath)

	// Open file and read its contents.
	fileContent, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open file: %v\n", err)
		return
	}

	fileAsBytes := make([]byte, 1000)
	count, err := fileContent.Read(fileAsBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read file: %v\n", err)
		return
	}

	if count == 0 {
		fmt.Fprintf(os.Stderr, "file is empty\n")
		return
	}

	fileAsBytes = fileAsBytes[:count]

	// src is the input that we want to tokenize.
	src := []byte(fileAsBytes)

	// sanitise the input

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
