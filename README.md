# go-reloaded

## Description

This project is a text correction and formatting tool written in Go.

It reads text from an input file, applies the required transformations, and writes the result into an output file.

## How to run

```bash
go run . input.txt output.txt
```

### Example

```bash
go run . sample.txt result.txt
cat result.txt
```

## Main features

The program can:

- convert `(hex)` values
- convert `(bin)` values
- apply `(up)`, `(low)`, and `(cap)`
- apply commands with numbers like `(up, 2)`
- fix punctuation spacing
- fix single quote spacing
- fix `a` / `an`

## Project structure

- `main.go` — handles arguments, file reading, and file writing
- `tokenize.go` — splits the text into tokens
- `commands.go` — applies commands to tokens
- `convert.go` — contains number conversion helpers
- `case.go` — contains manual case conversion functions
- `articles.go` — fixes `a` and `an`
- `build.go` — rebuilds the final text
- `helpers.go` — contains helper functions

## Error messages

If the number of arguments is wrong:

```text
Usage: go run . input.txt output.txt
```

If the input file cannot be read:

```text
Error reading input file
```

If the output file cannot be written:

```text
Error writing output file
```

## Notes

The program follows this order:

1. read the input file
2. tokenize the text
3. apply commands
4. fix articles
5. rebuild the final text
6. write the output file

For the full project behavior and detailed rules, see `PRD.md`.
