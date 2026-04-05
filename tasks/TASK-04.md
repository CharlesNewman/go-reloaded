# TASK-04 — Rebuild the final text with correct spacing

## Goal

Turn the processed tokens back into one clean text string with correct spacing around punctuation, quotes, and new lines.

## Why

After tokenizing and applying commands, the text is still stored as separate tokens.  
The program needs to rebuild those tokens into a final readable sentence or paragraph.

## Files involved

- `build.go`
- `helpers.go`

## What this part does

This part of the program goes through the final token list and joins everything back into one string.

While rebuilding the text, it must:

- attach punctuation to the previous word
- add a space after punctuation when needed
- keep grouped punctuation like `...` and `!?` together
- place single quotes correctly around the quoted text
- keep new lines in the final output

## Execution order in this task

This task runs after the program has already:

1. tokenized the text
2. applied commands
3. fixed `a/an`

Flow:

1. `main.go` receives the updated token list
2. `funcs.BuildText(tokens)` is called
3. each token is checked in order
4. spaces are added or removed depending on the token
5. the final text string is returned
6. `main.go` later writes this text into the output file

## Steps

1. Start with an empty result string.
2. Read tokens one by one.
3. If the token is a quote, place it correctly depending on whether the program is inside or outside quotes.
4. If the token is a new line, remove extra space before it and add the line break.
5. If the token is punctuation, remove any extra space before it and attach it to the previous word.
6. If the token is a normal word, add a space before it only when needed.
7. Continue until all tokens are rebuilt into one final string.

## Acceptance criteria

- Punctuation is attached to the previous word.
- There is a space after punctuation when needed.
- There is no extra space before punctuation.
- `...` stays together.
- `!?` stays together.
- Quotes are attached correctly around one word or many words.
- New lines stay in the final output.
- The final text is clean and readable.

## Tests

### Test 1 — Punctuation spacing

**Input tokens:**
```text
["I", "was", "there", ",", "and", "then", "BAMM", "!!"]
```

**Expected result:**
```text
I was there, and then BAMM!!
```

### Test 2 — Grouped punctuation

**Input tokens:**
```text
["I", "was", "thinking", "...", "You", "were", "right"]
```

**Expected result:**
```text
I was thinking... You were right
```

### Test 3 — One word in quotes

**Input tokens:**
```text
["I", "am", ":", "'", "awesome", "'"]
```

**Expected result:**
```text
I am: 'awesome'
```

### Test 4 — Many words in quotes

**Input tokens:**
```text
["As", "she", "said", ":", "'", "I", "will", "be", "there", "soon", "'"]
```

**Expected result:**
```text
As she said: 'I will be there soon'
```

### Test 5 — New line

**Input tokens:**
```text
["Hello", ",", "\n", "world", "!"]
```

**Expected result:**
```text
Hello,
world!
```

## Notes

- This task does not decide what words should change. It only rebuilds the final clean text.
- Quote handling in this task uses a small state idea: the program tracks whether it is inside quotes or not.
- This task is important because even correct transformations can look wrong if the final spacing is bad.
