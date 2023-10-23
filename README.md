# Grepclone

Grepclone is a simple Go application that allows you to search for a specific text string in all files within a given directory and its subdirectories. It utilizes a worklist-based concurrent approach to efficiently search for the target text across multiple files.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Package Structure](#package-structure)
- [Documentation](#documentation)
- [Contributing](#contributing)

## Installation

To use Grepclone, you need to have Go installed on your system. Follow these steps to install and build the application:

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/grepclone.git
   
2. Change to project directory:
   ```shell
   cd grepclone

3. Build the grepclone binary:
   ```shell
   go build -o grepclone main.go

This will create the `grepclone` binary in your current directory, which you can use to search for text in files.

## Usage

To use Grepclone, follow these steps:

1. Run the `grepclone` binary with the following command:
   
   ```shell
   ./grepclone <search_string> <directory_path>
   
  Replace `<search_string>` with the text you want to search for and `<directory_path>` with the path to the directory where you want to search for the text.

2. Grepclone will recursively search for the specified text within all files in the specified directory and its subdirectories.

3. Grepclone will display the matching lines along with their line numbers and file paths.

## Package Structure

Grepclone is organized into three packages: worklist, worker, and the main package.

### `worklist` Package
- Entry struct: Represents a file path in the worklist.
- Worklist struct: Manages a list of files to be processed concurrently.
  - Add(work Entry): Adds a file path to the worklist.
  - Next() Entry: Retrieves the next file path from the worklist.
  - New(bufSize int) Worklist: Creates a new worklist with the specified buffer size.
  - Finalize(numWorkers int): Terminates workers by adding empty paths to the worklist.
 
### `worker` Package
- Result struct: Represents a matching line in a file.
- Results struct: Stores multiple matching results.
- NewResult(line string, lineNum int, path string) Result: Creates a new result.
- FindInFile(path string, find string) *Results: Searches for the specified text in a file and returns matching results.
  
### Main Package
The main package (main) is responsible for the application's entry point, managing worker goroutines, and displaying search results.

## Documentation
### `worklist` Package

`Entry` Struct
Represents a file path in the worklist.
```go
type Entry struct {
 Path string
}
```

`Worklist` Struct
Manages a list of files to be processed concurrently.
```go
type Worklist struct {
 jobs chan Entry
}
```
### Methods
- `Add(work Entry)`: Adds a file path to the worklist.
- `Next() Entry`: Retrieves the next file path from the worklist.
- `New(bufSize int) Worklist`: Creates a new worklist with the specified buffer size.
- `Finalize(numWorkers int)`: Terminates workers by adding empty paths to the worklist.

### `worker` Package

`Result` Struct
Represents a matching line in a file.
```go
type Result struct {
 Line    string
 LineNum int
 Path    string
}
```

`Results` Struct
Stores multiple matching results.
```go
type Results struct {
 Inner []Result
}
```

### Methods
- NewResult(line string, lineNum int, path string) Result: Creates a new result.
- FindInFile(path string, find string) *Results: Searches for the specified text in a file and returns matching results.
  
### Main Package
The main package (main) contains the application's entry point, which manages the concurrent search for text in files within a specified directory.

### Contributing
Feel free to contribute to this project by opening issues, suggesting improvements, or submitting pull requests. Your feedback and contributions are highly appreciated.
