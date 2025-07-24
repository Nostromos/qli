---
sidebar_position: 2
---

# API Documentation

## Overview

The quiz application is structured as a single-file Go program that implements a command-line quiz game with timer functionality and CSV-based questions.

## Core Components

### Data Structures

#### Question
```go
type Question struct {
    prompt string
    answer string
}
```
Represents a single quiz question with its prompt and expected answer.

### Main Functions

#### `main()`
Entry point of the application. Handles:
- Command-line flag parsing
- Question loading
- User interaction flow
- Score display

#### `loadQuestions(filepath string) ([]Question, error)`
Loads questions from a CSV file.

**Parameters:**
- `filepath`: Path to the CSV file containing questions

**Returns:**
- `[]Question`: Slice of parsed questions
- `error`: Error if file cannot be read or parsed

**Error Handling:**
- Returns error if file doesn't exist
- Returns error if CSV format is invalid
- Returns error if file is empty

#### `shuffleQuestions(questions []Question)`
Shuffles questions in-place using Fisher-Yates algorithm.

**Parameters:**
- `questions`: Slice of questions to shuffle (modified in-place)

#### `runQuiz(questions []Question, limit time.Duration) int`
Main quiz execution logic. Manages question presentation, answer collection, and timing.

**Parameters:**
- `questions`: Slice of questions to present
- `limit`: Time duration for quiz completion

**Returns:**
- `int`: Number of correct answers

**Behavior:**
- Uses goroutines for concurrent timer and input handling
- Ends quiz when timer expires or all questions answered
- Returns immediately on timer expiration

#### `cleanInput(input string) string`
Normalizes user input for consistent answer matching.

**Parameters:**
- `input`: Raw user input string

**Returns:**
- `string`: Cleaned input (lowercase, trimmed whitespace)

## Concurrency Model

The application uses Go's concurrency features to handle simultaneous timer and user input:

```go
answerCh := make(chan string)  // Channel for user answers
doneCh := make(chan bool)      // Channel to signal quiz completion
```

### Timer Goroutine
- Runs in background during quiz
- Sends signal on `timer.C` when time expires
- Can be stopped early via `doneCh`

### Input Goroutine
- Reads user input in separate goroutine
- Prevents blocking on `scanner.Scan()`
- Sends answers through `answerCh`

## Command-Line Interface

### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `csv` | string | `"./questions.csv"` | Path to CSV file |
| `limit` | duration | `30s` | Time limit for quiz |
| `shuffle` | bool | `false` | Enable question shuffling |

### Flag Parsing
```go
flag.StringVar(&csvFilePath, "csv", "./questions.csv", "path to csv file")
flag.DurationVar(&limit, "limit", 30*time.Second, "quiz time limit")
flag.BoolVar(&shuffle, "shuffle", false, "shuffle questions")
```

## Error Handling Strategy

1. **File Errors**: Logged and cause program exit
2. **CSV Parsing**: Validated during load, exits on malformed data
3. **User Input**: Gracefully handled, invalid answers counted as incorrect
4. **Timer Issues**: Handled via select statement with multiple cases

## Testing Approach

The test suite (`quiz_test.go`) covers:

### File Operations
- Valid CSV loading
- Invalid file paths
- Empty files
- Malformed CSV data

### Question Shuffling
- Length preservation
- Order modification
- Randomization verification

### Input Processing
- Case normalization
- Whitespace handling
- Special character processing

### Timer Functionality
- Start/stop behavior
- Expiration handling
- Concurrent operation

## Best Practices Implemented

1. **Separation of Concerns**: Each function has a single, clear responsibility
2. **Error Propagation**: Errors are returned up the call stack
3. **Resource Management**: Channels are properly closed
4. **Concurrent Safety**: No shared state between goroutines
5. **Input Validation**: All user input is sanitized before use
6. **Graceful Shutdown**: Quiz ends cleanly on timer expiration