package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Represents a single question
type Question struct {
	prompt string
	answer string
}

// Represents a file containing a list of questions
type Quiz struct {
	path      string
	filetype  string
	questions [][]Question
}

// Represents a single game
type Game struct {
	quiz           Quiz
	correctAnswers int
	totalQuestions int
}

// ANSI colors for the live timer
// TODO: Refactor this out into a separate file along with all the CLI display logic
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// Default values for flags
const (
	defaultCSV = "./questions.csv"
	defaultTimeLimit = 30 * time.Second
)

func main() {
	// Define flags
	csvFilePath := flag.String("csv", defaultCSV, "A CSV file in the format of 'question,answer'")
	shuffleFlag := flag.Bool("shuffle", false, "Shuffle the quiz questions")
	timeLimit := flag.Duration("limit", defaultTimeLimit, "Time limit for the quiz - default is 30 seconds")
	flag.Parse()

	// Load questions from CSV file
	questions, err := loadQuestions(*csvFilePath)
	if err != nil {
		log.Fatalf("Failed to load questions: %v", err)
	}

	// OPTIONAL: Shuffle if flagged
	if *shuffleFlag {
		shuffleQuestions(questions)
	}


	// Welcome
	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("You have %v to answer as many questions as you can. Good luck!", *timeLimit)
	fmt.Println("Press Enter to start the quiz...")

	// Run quiz with time limit
	correctAnswers := runQuiz(questions, *timeLimit)
	fmt.Printf("You got %d out of %d questions correct!\n", correctAnswers, len(questions))
}

func loadQuestions(path string) ([]Question, error) {
	file, err := os.Open(strings.TrimSpace(path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	qs := make([]Question, len(records))
	for i, rec := range records {
		qs[i] = Question{
			prompt: rec[0], 
			answer: cleanInput(rec[1]), // TODO: Ensure that answers are trimmed and lowercased
		}
	}
	return qs, nil
}

func cleanInput(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func shuffleQuestions(questions []Question) {
	for i := range questions {
		j := rand.Intn(i + 1)
		questions[i], questions[j] = questions[j], questions[i]
	}
}

// func runTimer(start time.Time, line int, stop chan bool, color string) {
// 	for {
// 		select {
// 		case <-stop:
// 			return // Stop the timer when signaled
// 		default:
// 			elapsed := time.Since(start)
// 			// Use ANSI escape code to move cursor and update the correct line with color
// 			fmt.Printf("\033[%d;0H%sElapsed time for Timer %d: %s%s", line, color, line, elapsed.Round(time.Second), Reset)
// 			time.Sleep(100 * time.Millisecond) // Update every 100ms
// 		}
// 	}
// }
