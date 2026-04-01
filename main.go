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

// Order:
// 1. main.go
// 2. tokenize.go
// 3. applyCommands.go
// 4. doCommand.go
// 5. readCommand.go
// 6. hexToDecimalString.go
// 7. binToDecimalString.go
// 8. intToString.go
// 9. stringToInt.go
// 10. toUpperManual.go
// 11. toLowerManual.go
// 12. capitalizeManual.go
// 13. fixArticles.go
// 14. nextWord.go
// 15. startsWithVowelOrH.go
// 16. buildText.go
// 17. removeLastSpace.go
// 18. isCommand.go
// 19. isWord.go
// 20. isPunctuationChar.go
// 21. isPunctuationToken.go
// 22. findPreviousWord.go

// Program flow to study:
// read file -> tokenize -> apply commands -> fix articles -> build final text -> write file
