# go-reloaded

## Description
This program reads text from an input file, applies transformation commands written inside the text, fixes the article `a/an`, rebuilds the final text with correct spacing and punctuation, and writes the result to an output file.

---

## PRD

### Problem
The goal of this tool is to take a text file that contains normal text plus inline commands like `(up)`, `(low)`, `(cap)`, `(hex)`, and `(bin)`, then produce a clean transformed output file.

### Starting Plan
The starting idea of the project was:

1. Read text from a file.
2. Detect words, punctuation, quotes, commands, and line breaks.
3. Apply each command to the correct previous word or words.
4. Fix small grammar details like `a` / `an`.
5. Rebuild the text so punctuation spacing looks natural.
6. Save the final result into a new file.

### What the Final Code Does
The final program follows this flow:

- checks that exactly 2 arguments are given:
  - input file
  - output file
- reads the input file
- splits the text into tokens
- applies all supported commands
- fixes articles (`a` / `an`) and the oposit (`an` / `a`)
- rebuilds the final text with correct spacing
- writes the result to the output file

### Supported Transformations

#### 1. Number conversions
- `(hex)` converts the previous word from hexadecimal to decimal
- `(bin)` converts the previous word from binary to decimal

Example:
- `1E (hex)` → `30`
- `10 (bin)` → `2`

#### 2. Case changes
- `(up)` makes the previous word uppercase
- `(low)` makes the previous word lowercase
- `(cap)` capitalizes the previous word

Example:
- `go (up)` → `GO`
- `HELLO (low)` → `hello`
- `hELLo (cap)` → `Hello`

#### 3. Multi-word case changes
- `(up, N)`
- `(low, N)`
- `(cap, N)`

These apply to the previous `N` words.

Example:
- `this is nice (up, 2)` → `this IS NICE`

#### 4. Article fixing
The program checks `a`, `A`, `an`, and `An`, then changes them depending on the next word.

Example:
- `a apple` → `an apple`
- `an car` → `a car`

### Final Result
The code ended up as a small text-processing pipeline:

`read file -> tokenize -> apply commands -> fix articles -> rebuild text -> write file`

This makes the code easier to understand and keeps each job separated into its own function.

---

## CLI Analysis

### Command
```bash
go run . input.txt output.txt
```
### Arguments
- `input.txt` → file to read from
- `output.txt` → file to write the transformed result into

### Behavior

#### If arguments are missing
The program prints:

`Usage: go run . input.txt output.txt`

#### If the input file cannot be read
The program prints:

`Error reading input file`

#### If the output file cannot be written
The program prints:

`Error writing output file`
