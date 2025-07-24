# QLI - Quiz Line Interface

A command-line quiz application written in Go that reads questions from CSV files and challenges users against a timer.

## Features

- **CSV-based Questions**: Load quiz questions from CSV files
- **Timed Quizzes**: Set custom time limits for completing the quiz
- **Question Shuffling**: Randomize question order for variety
- **Clean Interface**: Simple, intuitive command-line interface
- **Score Tracking**: Displays final score upon completion
- **Flexible Answer Matching**: Case-insensitive with whitespace normalization

## Requirements

- Go 1.23.0 or higher

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd quiz
```

2. Build the application:
```bash
go build -o quiz quiz.go
```

Or run directly without building:
```bash
go run quiz.go
```

## Usage

### Basic Usage

Run the quiz with default settings (30-second timer, questions.csv file):
```bash
./quiz
```

### Command-Line Options

```bash
./quiz [options]
```

| Flag | Description | Default |
|------|-------------|---------|
| `-csv` | Path to CSV file containing questions | `./questions.csv` |
| `-limit` | Time limit for the quiz (e.g., 30s, 1m, 1h) | `30s` |
| `-shuffle` | Shuffle questions randomly | `false` |

### Examples

```bash
# Use a custom questions file
./quiz -csv=myquestions.csv

# Set a 1-minute time limit
./quiz -limit=1m

# Shuffle questions for variety
./quiz -shuffle

# Combine all options
./quiz -csv=advanced.csv -limit=2m -shuffle
```

## CSV File Format

Questions must be in a CSV file with exactly 2 columns:
- **Column 1**: Question text
- **Column 2**: Answer

### Example CSV:
```csv
5+5,10
What is the capital of France?,paris
What year was Go released?,2009
```

**Note**: The CSV parser properly handles questions containing commas.

## How It Works

1. **Start**: Press Enter to begin the quiz and start the timer
2. **Questions**: Each question is displayed one at a time
3. **Answers**: Type your answer and press Enter to continue
4. **Timer**: The quiz ends when time runs out or all questions are answered
5. **Score**: Final score is displayed as "X out of Y questions correct"

## Answer Matching

Answers are matched with the following rules:
- Case-insensitive (e.g., "Paris" matches "paris")
- Leading/trailing whitespace is ignored
- Exact match required after normalization

## Development

### Running Tests

```bash
go test
```

### Project Structure

```
quiz/
├── quiz.go          # Main application code
├── quiz_test.go     # Test suite
├── questions.csv    # Default questions file
├── go.mod          # Go module file
└── README.md       # This file
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is part of the [Gophercises](https://gophercises.com/) exercise series by Jon Calhoun.

## Acknowledgments

- Built as part of the Gophercises learning series
- Focuses on teaching Go fundamentals including:
  - File I/O
  - CSV parsing
  - Command-line flags
  - Concurrency with goroutines and channels
  - Timer management
  - Testing practices