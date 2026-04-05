# TASK-03 — Apply text transformation commands

## Goal

Apply the special commands found in the token list so words are converted correctly before the final text is rebuilt.

## Why

The main feature of the project is changing text based on commands written inside the input.  
If this task is wrong, the program will keep the commands in the text or produce the wrong result.

## Files involved

- `commands.go`
- `convert.go`
- `case.go`

## What this part does

This part of the program looks through the token list and checks if a token is a command.

If the token is a command, it changes the previous word or previous words depending on the command.

The supported commands are:

- `(hex)` → convert the previous hexadecimal word to decimal
- `(bin)` → convert the previous binary word to decimal
- `(up)` → make the previous word uppercase
- `(low)` → make the previous word lowercase
- `(cap)` → capitalize the previous word
- `(up, n)` → uppercase the previous `n` words
- `(low, n)` → lowercase the previous `n` words
- `(cap, n)` → capitalize the previous `n` words

## Execution order in this task

This task runs after tokenizing.

Flow:

1. `main.go` receives the tokens from `Tokenize`
2. `funcs.ApplyCommands(tokens)` is called
3. each token is checked
4. command tokens are read and understood
5. the correct previous word or words are changed
6. command tokens are not kept in the final list
7. the updated token list is returned

## Steps

1. Loop through all tokens.
2. Check if the current token is a command.
3. If it is not a command, keep it.
4. If it is a command, read its name and count.
5. Find the previous word or previous words.
6. Apply the correct transformation.
7. Skip the command token itself from the result.
8. Return the updated token list.

## Acceptance criteria

- `(hex)` changes the previous hexadecimal number to decimal.
- `(bin)` changes the previous binary number to decimal.
- `(up)` changes the previous word to uppercase.
- `(low)` changes the previous word to lowercase.
- `(cap)` capitalizes the previous word.
- `(up, n)` changes the previous `n` words to uppercase.
- `(low, n)` changes the previous `n` words to lowercase.
- `(cap, n)` capitalizes the previous `n` words.
- Command tokens do not stay in the final token list.
- If a value is not valid for conversion, it stays unchanged.

## Tests

### Test 1 — Hex conversion

**Input tokens:**
```text
["1E", "(hex)", "files", "were", "added"]
```

**Expected result:**
```text
["30", "files", "were", "added"]
```

### Test 2 — Binary conversion

**Input tokens:**
```text
["10", "(bin)", "years"]
```

**Expected result:**
```text
["2", "years"]
```

### Test 3 — Uppercase command

**Input tokens:**
```text
["go", "(up)"]
```

**Expected result:**
```text
["GO"]
```

### Test 4 — Lowercase command

**Input tokens:**
```text
["SHOUTING", "(low)"]
```

**Expected result:**
```text
["shouting"]
```

### Test 5 — Capitalize command

**Input tokens:**
```text
["brooklyn", "(cap)"]
```

**Expected result:**
```text
["Brooklyn"]
```

### Test 6 — Command with count

**Input tokens:**
```text
["so", "exciting", "(up, 2)"]
```

**Expected result:**
```text
["SO", "EXCITING"]
```

### Test 7 — Mixed command line

**Input tokens:**
```text
["Simply", "add", "42", "(hex)", "and", "10", "(bin)"]
```

**Expected result:**
```text
["Simply", "add", "66", "and", "2"]
```

## Notes

- This task changes tokens only. It does not rebuild the final spacing.
- `convert.go` helps with number conversion.
- `case.go` helps with uppercase, lowercase, and capitalize logic.
- This task is one of the most important parts of the whole project because it performs the main text transformations.
