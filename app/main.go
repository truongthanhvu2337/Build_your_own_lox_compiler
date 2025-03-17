package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/internal/errorutil"
	"github.com/codecrafters-io/interpreter-starter-go/internal/scanner"
)


func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}


	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	contents := string(fileContents)
	scan := scanner.NewScanner(contents)
	tokens := scan.ScanTokens()
	for _, token := range tokens {
        fmt.Printf("%s  %v\n", token.Type, token.Literal)
    }

	if errorutil.GlobalErrorHandler.HadError {
		os.Exit(65)
	}	
}
