# PRD — go-reloaded

## 1. Project Overview

**go-reloaded** is a text processing tool written in Go.
It reads a text file, applies formatting and correction rules, and writes the corrected result into another file.

## 2. Purpose

The purpose of this program is to automate simple text editing rules.

The tool must:

- read text from an input file
- detect special commands inside the text
- modify words based on those commands
- fix punctuation spacing
- fix single quote spacing
- fix `a` / `an`
- save the final result in an output file

## 3. Command Line Usage

The program is executed from the terminal like this:

```bash
go run . input.txt output.txt
```

### Arguments

- `input.txt` -> the file to read
- `output.txt` -> the file where the final text will be written

## 4. Input and Output

### Input

The input file contains text that may include:

- normal words
- punctuation with bad spacing
- quotes with bad spacing
- commands like `(up)` or `(hex)`

### Output

The output file contains the corrected version of the text after all rules are applied.

## 5. Required Features

### 5.1 Hexadecimal conversion

When the program finds `(hex)`, it converts the word before it from hexadecimal to decimal.

**Example:**

```text
1E (hex) files were added
```

**Output:**

```text
30 files were added
```

### 5.2 Binary conversion

When the program finds `(bin)`, it converts the word before it from binary to decimal.

**Example:**

```text
It has been 10 (bin) years
```

**Output:**

```text
It has been 2 years
```

### 5.3 Uppercase conversion

When the program finds `(up)`, it changes the previous word to uppercase.

**Example:**

```text
Ready, set, go (up) !
```

**Output:**

```text
Ready, set, GO!
```

### 5.4 Lowercase conversion

When the program finds `(low)`, it changes the previous word to lowercase.

**Example:**

```text
I should stop SHOUTING (low)
```

**Output:**

```text
I should stop shouting
```

### 5.5 Capitalize conversion

When the program finds `(cap)`, it capitalizes the previous word.

**Example:**

```text
Welcome to the brooklyn bridge (cap)
```

**Output:**

```text
Welcome to the brooklyn Bridge
```

### 5.6 Commands with numbers

The commands `(up, n)`, `(low, n)`, and `(cap, n)` apply to the previous `n` words.

**Example:**

```text
This is so exciting (up, 2)
```

**Output:**

```text
This is SO EXCITING
```

## 6. Punctuation Rules

The punctuation marks:

- `.`
- `,`
- `!`
- `?`
- `:`
- `;`

must be placed:

- directly after the previous word
- with one space after them when needed

**Example:**

```text
I was sitting over there ,and then BAMM !!
```

**Output:**

```text
I was sitting over there, and then BAMM!!
```

### Grouped punctuation

Grouped punctuation like `...` or `!?` must stay together.

**Example:**

```text
I was thinking ... You were right
```

**Output:**

```text
I was thinking... You were right
```

## 7. Quote Rules

Single quotes `'` must be attached correctly around the text inside them.

### One word inside quotes

**Example:**

```text
I am exactly how they describe me: ' awesome '
```

**Output:**

```text
I am exactly how they describe me: 'awesome'
```

### Many words inside quotes

**Example:**

```text
As she said: ' I will be there soon '
```

**Output:**

```text
As she said: 'I will be there soon'
```

## 8. Article Correction Rule

The program must change `a` into `an` when the next word begins with:

- a vowel: `a`, `e`, `i`, `o`, `u`
- `h`

**Example:**

```text
There it was. A amazing rock!
```

**Output:**

```text
There it was. An amazing rock!
```

### ***I also handle special cases***

**Example:**

```text
It was a honest mistake.
```

**Output:**

```text
It was an honest mistake.
```

## 9. Error Handling

The program must handle common errors clearly.

### Wrong number of arguments

If the user does not give exactly 2 file names, the program prints:

```text
Usage: go run . input.txt output.txt
```

### Input file cannot be read

If the input file does not exist or cannot be opened, the program prints:

```text
Error reading input file
```

### Output file cannot be written

If the output file cannot be created or written, the program prints:

```text
Error writing output file
```

## 10. Program Flow

The program follows these main steps:

1. check command-line arguments
2. read the input file
3. convert the file content into tokens
4. apply commands like `(up)` and `(hex)`
5. fix articles (`a` / `an`)
6. rebuild the final text with correct spacing
7. write the result into the output file

## 11. Code Structure

The project is separated into small files with clear jobs.

### `main.go`
Handles arguments, file reading, file writing, and the main program flow.

### `tokenize.go`
Splits the text into tokens like words, punctuation, quotes, commands, and new lines.

### `commands.go`
Detects and applies commands such as `(hex)`, `(bin)`, `(up)`, `(low)`, and `(cap)`.

### `convert.go`
Contains helper functions for number conversion and string-to-int / int-to-string logic.

### `case.go`
Contains manual uppercase, lowercase, and capitalize functions.

### `articles.go`
Fixes `a` and `an` depending on the next word.

### `build.go`
Rebuilds the final clean text from the tokens.

### `helpers.go`
Contains small helper functions for punctuation and special word rules.

## 12. Example

### Input

```text
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) .
```

### Output

```text
It was the best of times, it was the worst of TIMES, It Was The Age Of Foolishness.
```

## 13. Goal of the Project

This project helps practice:

- Go file handling
- string manipulation
- token processing
- rebuilding formatted text
- thinking step by step
- organizing code into small functions

## 14. Limitations and Known Edge Cases

This project handles the rules required by the subject, but some edge cases are not fully supported.

### Current limitations

- Single quotes are handled as quote markers, so apostrophes inside words like `don't` or `I'm` are not treated separately.
- If quotes are not balanced correctly in the input, the final spacing may not be perfect.
- The `a/an` rule handles the main required cases and some common sound exceptions, but it does not cover every possible English pronunciation case.
- The manual case conversion functions are designed for basic English letters and are not intended for full Unicode text handling.
- Invalid `(hex)` or `(bin)` values are left unchanged instead of producing an error message.

### Note

These limitations do not affect the main required behavior of the project examples from the subject, but they are important to keep in mind as edge cases.