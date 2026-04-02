# TASK-01 — Handle CLI arguments and file I/O

## Goal

Make the program start correctly from the terminal, accept the input and output file names, read the input file, and write the final result into the output file.

## Why

This task is the entry point of the whole project.  
If the program cannot receive the file names or cannot read and write files correctly, the rest of the logic cannot run.

## Files involved

- `main.go`

## What this part does

This part of the program:

1. checks that exactly 2 arguments are given
2. saves the input file name
3. saves the output file name
4. reads the input file
5. sends the text to the processing steps
6. writes the final result into the output file
7. prints clear error messages when something fails

## Execution order in this task

Inside `main.go`, the flow is:

1. check `os.Args`
2. read the input file with `os.ReadFile`
3. call:
   - `funcs.Tokenize`
   - `funcs.ApplyCommands`
   - `funcs.FixArticles`
   - `funcs.BuildText`
4. add a final newline
5. write the result with `os.WriteFile`

## Steps

1. Check that the user gave exactly 2 file names.
2. Print the usage message if the number of arguments is wrong.
3. Read the input file.
4. Print an error if the input file cannot be read.
5. Convert the file data into a string.
6. Pass the text through the project functions in the correct order.
7. Add a newline at the end of the result.
8. Write the final text to the output file.
9. Print an error if the output file cannot be written.

## Acceptance criteria

- The program runs with:
  ```bash
  go run . input.txt output.txt
  ```
- If arguments are missing, it prints:
  ```text
  Usage: go run . input.txt output.txt
  ```
- If the input file cannot be read, it prints:
  ```text
  Error reading input file
  ```
- If the output file cannot be written, it prints:
  ```text
  Error writing output file
  ```
- If everything works, the program creates the output file with the processed text.

## Tests

### Test 1 — Missing arguments

**Command:**
```bash
go run .
```

**Expected result:**
```text
Usage: go run . input.txt output.txt
```

### Test 2 — Missing input file

**Command:**
```bash
go run . missing.txt result.txt
```

**Expected result:**
```text
Error reading input file
```

### Test 3 — Valid files

**Command:**
```bash
go run . sample.txt result.txt
```

**Expected result:**

- the program runs without crashing
- `result.txt` is created
- the output contains the processed text

## Notes

- This task does not explain the inside logic of tokenizing or applying commands.
- Its job is to start the program and connect all the processing parts together.
- This is the file that controls the full execution order of the project.
