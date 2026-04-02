# AI Task Decomposition Index — go-reloaded

## Purpose

This file lists the small task cards created for the project and shows how the work was split into testable parts.

The project was divided into 5 tasks so that each important part of the pipeline could be understood, built, and checked step by step.

## Why this decomposition was chosen

The program follows a pipeline architecture:

1. read the input file
2. tokenize the text
3. apply transformation commands
4. fix articles
5. rebuild the final text
6. write the output file

Because of that, the tasks were grouped around the main stages of the program.

This makes the work easier to explain during an audit and easier to test in small pieces.

## Task list

### TASK-01 — Handle CLI arguments and file I/O
**Focus:** start the program, read the input file, call the processing steps, and write the output file.

**Files involved:**
- `main.go`

**Why it comes first:**
The program cannot do anything until it receives the correct arguments and reads the input file.

---

### TASK-02 — Tokenize the input text
**Focus:** split the input text into tokens such as words, punctuation, quotes, commands, and new lines.

**Files involved:**
- `tokenize.go`

**Why it comes second:**
The program needs tokens before it can apply changes safely.

---

### TASK-03 — Apply text transformation commands
**Focus:** apply `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`, and command versions with numbers.

**Files involved:**
- `commands.go`
- `convert.go`
- `case.go`

**Why it comes third:**
Commands must be applied before the final text is rebuilt.

---

### TASK-04 — Rebuild the final text with correct spacing
**Focus:** rebuild the final output text with correct punctuation spacing, quote placement, grouped punctuation, and new lines.

**Files involved:**
- `build.go`
- `helpers.go`

**Why it comes fourth:**
After the text has been changed, it must be rebuilt into a clean readable string.

---

### TASK-05 — Fix articles and do final behavior checks
**Focus:** fix `a/an`, handle simple pronunciation exceptions, and confirm that the final behavior matches the subject rules.

**Files involved:**
- `articles.go`
- `helpers.go`

**Why it comes last:**
This is a final correction step before the text is fully finished.

## Dependency order

The tasks depend on each other in this order:

- `TASK-01` → prepares input and program flow
- `TASK-02` → creates tokens from the raw text
- `TASK-03` → changes tokens using commands
- `TASK-05` → fixes articles on the processed tokens
- `TASK-04` → rebuilds the final clean text

## Note about execution order

The project task order and the exact execution order are very close, but not identical.

The real execution order inside the program is:

1. `main.go`
2. `Tokenize`
3. `ApplyCommands`
4. `FixArticles`
5. `BuildText`
6. write output file

This means that `TASK-05` runs before `TASK-04` during execution, even if `TASK-04.md` was written first as a task card.

## Why 5 tasks are enough

The project instructions ask for 5–8 small testable tasks.

This decomposition uses 5 tasks because:

- each task has one clear purpose
- all project requirements are covered
- the files are grouped in a simple and believable way
- the project is small enough that more tasks are not necessary

## Final summary

The task decomposition covers the full go-reloaded project from input handling to final output generation.

It was designed to:

- match the real structure of the code
- stay simple and beginner-friendly
- make the project easier to explain during audit
- keep each task small enough to test
