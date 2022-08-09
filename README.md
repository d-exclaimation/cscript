# CScript

Compiler for C source code to both source and header file

## Installation

1. Install go
2. Run `go install`
3. Set `go/bin`n to path

## Usage

```shell
cscript <input> <output>
```

- `input`: Input filename with extensions
    - File can be either `.c`, `.h`, or `.cscript`
    - `.c` or `.h` will be replaced with the compiled ones unless specified a different `output`.
    - Recommend using `.cscript` to perform transpiling if the code is frequently updated, as it is not replaced
- `output`: Output filename without extensions (optional, defaults to `input`)

### Algorithms / Modes

- `auto-flag`, check based on the position of the first full non-`static` function.

### Limitation

Currently, the CLI can only run in `auto-flag` mode which judge declarations such as includes, defines, typedefs,
etc.
by checking the position of first function (with a executable block). This has limitation when declarations are
scattered
across the file.

It's is recommended to keep these declarations declared before any executable function.

## Why?

Apparently, this tool has not existed yet in an accessible manner. In short,
I am way too lazy to write a separate header file especially when it's clearly possible to automate the task.