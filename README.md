# CScript

Compiler for C source code to both source and header file

## Installation

1. Install go
2. Run `go install`
3. Set `go/bin`n to path

## Usage

```shell
cscript <filename>
```

- `filename`: File can be either `.c`, `.h`, or `.cscript`
    - `.c` or `.h` will be replaced with the compiled ones
    - Recommend using `.cscript` to perform transpiling if the code is frequently updated, as it is not replaced

## Why?

Apparently, this tool has not exist yet in a accessible manner. In short,
I am way to lazy to write a separate header file.