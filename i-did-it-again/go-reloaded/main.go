package main

import (
	"fmt"
	"go-reloaded/funcs"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . input.txt output.txt")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading the input file")
		return
	}
	text := string(data)
	fixedText := funcs.ProcessText(text)

	err = os.WriteFile(output, []byte(fixedText), 0644)
	if err != nil {
		fmt.Println("Error writing the output file")
		return
	}

}
