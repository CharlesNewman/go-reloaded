# Golden Tests — go-reloaded

## Purpose

This file contains manual golden tests for the **go-reloaded** project.
A golden test compares a fixed input with a fixed expected output.
If the program output matches the expected output exactly, the test passes.

These tests help check that the program behavior stays correct after changes.

---

## How to Use These Tests

For each test:

1. Put the **Input** text into a file such as `sample.txt`
2. Run the program:

```bash
go run . sample.txt result.txt
```

3. Compare the content of `result.txt` with the **Expected Output**
4. If they are exactly the same, the test passes

---

## Test 01 — `(hex)` conversion

### What it checks
Converts the previous hexadecimal word into decimal.

### Input
```text
1E (hex) files were added
```

### Expected Output
```text
30 files were added
```

---

## Test 02 — `(bin)` conversion

### What it checks
Converts the previous binary word into decimal.

### Input
```text
It has been 10 (bin) years
```

### Expected Output
```text
It has been 2 years
```

---

## Test 03 — `(up)` conversion

### What it checks
Changes the previous word to uppercase.

### Input
```text
Ready, set, go (up) !
```

### Expected Output
```text
Ready, set, GO!
```

---

## Test 04 — `(low)` conversion

### What it checks
Changes the previous word to lowercase.

### Input
```text
I should stop SHOUTING (low)
```

### Expected Output
```text
I should stop shouting
```

---

## Test 05 — `(cap)` conversion

### What it checks
Capitalizes the previous word.

### Input
```text
Welcome to the brooklyn bridge (cap)
```

### Expected Output
```text
Welcome to the brooklyn Bridge
```

---

## Test 06 — `(up, n)` conversion

### What it checks
Changes the previous `n` words to uppercase.

### Input
```text
This is so exciting (up, 2)
```

### Expected Output
```text
This is SO EXCITING
```

---

## Test 07 — punctuation spacing

### What it checks
Moves punctuation next to the previous word and leaves one space after it when needed.

### Input
```text
I was sitting over there ,and then BAMM !!
```

### Expected Output
```text
I was sitting over there, and then BAMM!!
```

---

## Test 08 — grouped punctuation

### What it checks
Keeps grouped punctuation together, like `...`.

### Input
```text
Punctuation tests are ... kinda boring ,what do you think ?
```

### Expected Output
```text
Punctuation tests are... kinda boring, what do you think?
```

---

## Test 09 — single quotes with one word

### What it checks
Removes extra spaces inside single quotes around one word.

### Input
```text
I am exactly how they describe me: ' awesome '
```

### Expected Output
```text
I am exactly how they describe me: 'awesome'
```

---

## Test 10 — single quotes with many words

### What it checks
Removes extra spaces inside single quotes around a full sentence.

### Input
```text
As she said: ' I will be there soon '
```

### Expected Output
```text
As she said: 'I will be there soon'
```

---

## Test 11 — article correction

### What it checks
Changes `a` to `an` before a vowel or `h`.

### Input
```text
There it was. A amazing rock!
```

### Expected Output
```text
There it was. An amazing rock!
```

## Test 12 — article correction special cases

### Input:

```text
It was a honest mistake.
```

### Expected Output

```text
It was an honest mistake.
```

---

## Test 13 — mixed full example

### What it checks
Tests several rules together in one input.

### Input
```text
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) .
```

### Expected Output
```text
It was the best of times, it was the worst of TIMES, It Was The Age Of Foolishness.
```

---

