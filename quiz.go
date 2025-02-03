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
	defaultCSV       = "./questions.csv"
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
	fmt.Printf("You have %v to answer as many questions as you can. Good luck!\n", *timeLimit)
	fmt.Println("Press Enter to start the quiz...")
	waitForEnter()

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
			answer: cleanInput(rec[1]),
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
	// Traverses through questions and swaps with a random index
	for i := range questions {
		j := rand.Intn(i + 1)
		questions[i], questions[j] = questions[j], questions[i]
	}
}

func waitForEnter() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

func runQuiz(questions []Question, timeLimit time.Duration) int {
	correctAnswers := 0
	timer := time.NewTimer(timeLimit)
	answerCh := make(chan string)
	reader := bufio.NewReader(os.Stdin)

quizLoop:
	for i, q := range questions {
		fmt.Printf("Question %d: %s =", i+1, q.prompt)

		// So goroutines were the only way I could think of to handle the user input block (calling reader.Readstring() blocks input until a user types an answer and I'm supposed to kill the quiz even while waiting for an answer). I'm not sure if this is the best way to do it.
		go func() {
			input, _ := reader.ReadString('\n')
			answerCh <- cleanInput(input)
		}()

		// Wait for answer or timeout
		select {
		case <-timer.C: // Per https://pkg.go.dev/time#Timer, when the timer expires, current time gets sent on on C, which is our signal to alert the user and break the loop (thus ending the quiz). The JS/TS dev in me hates breaking loops but it seems like go devs are ok with it, given their baser natures.
			fmt.Println("\nTime's up!")
			break quizLoop
		case answer := <-answerCh:
			if answer == q.answer {
				correctAnswers++
			}
		}
	}
	return correctAnswers
}
