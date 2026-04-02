# TASK-05 — Fix articles and do final behavior checks

## Goal

Fix `a` and `an` based on the next word, including simple sound-based exceptions, and verify that the full program behavior matches the project rules.

## Why

Even if the commands and spacing work, the final output can still look wrong if articles are incorrect.  
This task improves the natural reading of the text and helps confirm that the whole pipeline works correctly.

## Files involved

- `articles.go`
- `helpers.go`

## What this part does

This part of the program goes through the tokens and checks words like:

- `a`
- `A`
- `an`
- `An`

It looks at the next real word and decides whether the article should be:

- `a`
- `an`
- `A`
- `An`

The code also handles some special sound cases:

- silent `h` words like `honest`, `honor`, and `hour`
- words that begin with vowel letters but sound like `you`, such as `user`, `university`, `European`, and `one`

## Execution order in this task

This task runs after commands have already been applied and before the final text is rebuilt.

Flow:

1. `main.go` receives the updated token list from `ApplyCommands`
2. `funcs.FixArticles(tokens)` is called
3. each article token is checked
4. the next valid word is found
5. the article is corrected if needed
6. the updated token list is returned
7. later, `BuildText` turns the tokens into the final output string

## Steps

1. Loop through the tokens.
2. Find tokens that are `a`, `A`, `an`, or `An`.
3. Find the next real word after the article.
4. Ignore punctuation and commands while searching for that next word.
5. Check whether the next word should use `a` or `an`.
6. Apply the correct lowercase or uppercase version.
7. Keep going until all tokens are checked.
8. Use final example cases to confirm the full program works as expected.

## Acceptance criteria

- `a` changes to `an` before words that begin with vowel sounds.
- `A` changes to `An` before words that begin with vowel sounds.
- `an` changes back to `a` when the next word should not use `an`.
- silent `h` words are handled correctly.
- words with a `you` sound are handled correctly.
- punctuation between words does not break the rule.
- the corrected tokens can still be rebuilt cleanly by the next step.

## Tests

### Test 1 — Basic vowel case

**Input tokens:**
```text
["a", "apple"]
```

**Expected result:**
```text
["an", "apple"]
```

### Test 2 — Uppercase article

**Input tokens:**
```text
["A", "amazing", "rock"]
```

**Expected result:**
```text
["An", "amazing", "rock"]
```

### Test 3 — Silent h case

**Input tokens:**
```text
["a", "honest", "mistake"]
```

**Expected result:**
```text
["an", "honest", "mistake"]
```

### Test 4 — You-sound case

**Input tokens:**
```text
["an", "user"]
```

**Expected result:**
```text
["a", "user"]
```

### Test 5 — Another you-sound case

**Input tokens:**
```text
["an", "European", "city"]
```

**Expected result:**
```text
["a", "European", "city"]
```

### Test 6 — End-to-end behavior check

**Input text:**
```text
There it was. A amazing rock!
```

**Expected final output:**
```text
There it was. An amazing rock!
```

## Notes

- This task improves correctness after the main text transformations are done.
- The rule is based on sound in some common cases, not only the first letter.
- This task is small, but it makes the final output feel much more natural and polished.
