# Tiny Search Engine

## Description

Tiny Search Engine is a simple Go program that searches for a word or phrase in a text file and prints the line numbers where the search term is found.

This project demonstrates basic file handling, user input, string searching, and looping in Go.

---

## Features

- Reads text from a file
- Accepts a search term from the user
- Searches each line of the file
- Prints the line numbers containing the search term
- Handles file opening and reading errors

---

## Project Structure

```
.
├── main.go
├── text.txt
└── README.md
```

---

## Example

### `text.txt`

```
Go is fast.
Go is simple.
Python is popular.
```

### Input

```
Search: Go
```

### Output

```
Found on line 1
Found on line 2
```

---

## How It Works

1. Opens `text.txt`.
2. Prompts the user for a search term.
3. Reads the file one line at a time.
4. Checks whether each line contains the search term.
5. Prints the corresponding line numbers for matches.

---

## Concepts Practiced

- Variables
- User input (`fmt.Scanln`)
- File handling (`os.Open`)
- Reading files with `bufio.Scanner`
- Loops
- Conditional statements
- String searching (`strings.Contains`)
- Error handling
- Line counting

---

## Run the Project

```bash
go run .
```

or

```bash
go run main.go
```

---

## Future Improvements

- Case-insensitive search
- Whole-word matching
- Display matching lines
- Count total matches
- Accept the filename and search term as command-line arguments