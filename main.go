package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] 
	if len(args) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]",
			"\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}
	outputFile := ""
	if strings.HasPrefix(args[0], "--") {
		flag := args[0]
		if !strings.HasPrefix(flag, "--output=") || len(flag) <= len("--output=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]",
				"\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(1)
		}

		outputFile = strings.TrimPrefix(flag, "--output=") 
		args = args[1:]                                    
	}
	if len(args) == 0 || len(args) > 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]",
			"\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}

	text := args[0]
	banner := "standard" 
	if len(args) == 2 {
		banner = args[1]
	}
	result, err := GenerateASCII(text, banner)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if outputFile != "" {
		err = os.WriteFile(outputFile, []byte(result), 0o644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(result)
	}
}
