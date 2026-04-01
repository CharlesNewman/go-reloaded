package main

import (
	"fmt"
	"go-reloaded/funcs"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file")
		return
	}

	text := string(data)

	tokens := funcs.Tokenize(text)
	tokens = funcs.ApplyCommands(tokens)
	tokens = funcs.FixArticles(tokens)
	result := funcs.BuildText(tokens)

	result = result + "\n"

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output file")
		return
	}
}

// os.WriteFile needs 3 arguments: filename, file data, file permissions.
// The number 0644 is the permission 6:
//0644
//│││└─ others = 4 = read
//││└── group  = 4 = read
//│└─── owner  = 6 = read + write
//└──── octal (this tells go this number is octal, not normal decimal)

//Order:
// 1. main.go
// 2. funcs/tokenize.go
// 3. funcs/commands.go
// 4. funcs/convert.go
// 5. funcs/case.go
// 6. funcs/articles.go
// 7. funcs/build.go
// 8. funcs/helpers.go
