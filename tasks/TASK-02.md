# TASK-02 — Tokenize the input text

## Goal

Split the input text into small parts called tokens so the rest of the program can process the text step by step.

## Why

The program cannot apply commands or rebuild the text correctly if the input stays as one long string.  
Tokenizing makes the text easier to understand and easier to change.

## Files involved

- `tokenize.go`

## What this part does

This part of the program reads the full text character by character and splits it into tokens.

The tokens can be:

- normal words
- punctuation marks
- grouped punctuation like `...`
- grouped punctuation like `!!` or `!?`
- single quotes `'`
- commands like `(up)` or `(cap, 2)`
- new line markers

## Execution order in this task

This task runs after the input file has been read in `main.go`.

Flow:

1. `main.go` turns the file data into a string
2. `funcs.Tokenize(text)` is called
3. the text is split into tokens
4. the token list is returned to `main.go`
5. the next step uses those tokens

## Steps

1. Start with an empty list of tokens.
2. Read the text one character at a time.
3. Build normal words until a separator is found.
4. Save words when a space or tab is reached.
5. Save `\n` as a separate token for new lines.
6. Detect commands starting with `(` and ending with `)`.
7. Save punctuation as separate tokens.
8. Keep `...` together as one token.
9. Keep groups like `!!` or `!?` together as one token.
10. Save single quotes `'` as separate tokens.

## Acceptance criteria

- Words are split correctly into separate tokens.
- Commands like `(up)` stay as one token.
- Commands like `(cap, 2)` stay as one token.
- Punctuation marks like `,` and `!` become separate tokens.
- `...` stays together as one token.
- `!!` or `!?` stays together as one token.
- Single quotes are separated correctly.
- New lines are kept as `\n` tokens.

## Tests

### Test 1 — Simple words

**Input:**
```text
hello world
```

**Expected tokens:**
```text
["hello", "world"]
```

### Test 2 — Command token

**Input:**
```text
go (up)
```

**Expected tokens:**
```text
["go", "(up)"]
```

### Test 3 — Command with number

**Input:**
```text
so exciting (up, 2)
```

**Expected tokens:**
```text
["so", "exciting", "(up, 2)"]
```

### Test 4 — Punctuation

**Input:**
```text
Hello , world !
```

**Expected tokens:**
```text
["Hello", ",", "world", "!"]
```

### Test 5 — Grouped punctuation

**Input:**
```text
Wait ... what !?
```

**Expected tokens:**
```text
["Wait", "...", "what", "!?"]
```

### Test 6 — Quotes

**Input:**
```text
' awesome '
```

**Expected tokens:**
```text
["'", "awesome", "'"]
```

## Notes

- This task only splits the text. It does not apply any transformations yet.
- Good tokenizing makes the next tasks much easier.
- This file is one of the base parts of the whole project because all later steps depend on correct tokens.
