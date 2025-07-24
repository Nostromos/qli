package main

import (
	"os"
	"testing"
	"time"
)

func createTempFile(t *testing.T, content string) string {
	t.Helper() // Marks this as a test helper function
	tmpfile, err := os.CreateTemp("", "quiz_test*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		tmpfile.Close()
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}
	return tmpfile.Name()
}

// Quiz File Loading
// 1. Must load questions correctly from a valid CSV file.
// 2. Must properly handle invalid/malformed CSV files.
// 3. Must properly handle empty CSV files.

// TestLoadQuestions_ValidFile // test that questions from a valid CSV are loaded correctly
func TestLoadQuestions_ValidFile(t *testing.T) {
	content := "What is 2 + 2?,4\nWhat is the capital of France?,paris"
	path := createTempFile(t, content)
	defer os.Remove(path) // Cleanup

	questions, err := loadQuestions(path)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(questions) != 2 {
		t.Fatalf("Expected 2 questions, got %d", len(questions))
	}

	if questions[0].prompt != "What is 2 + 2?" || questions[0].answer != "4" {
		t.Fatalf("Unexpected question/answer pair: %v", questions[0])
	}
}

// TestLoadQuestions_InvalidFile // test that an error is thrown when loading an invalid path
func TestLoadQuestions_InvalidFile(t *testing.T) {
	_, err := loadQuestions("nonexistent.csv")
	if err == nil {
		t.Fatal("Expected an error when loading a non-existent file, got nil")
	}
}

// TestLoadQuestions_EmptyFile // test that an error is thrown when loading an empty file
func TestLoadQuestions_EmptyFile(t *testing.T) {
	path := createTempFile(t, "")
	defer os.Remove(path)

	questions, err := loadQuestions(path)
	if err == nil {
		t.Fatal("Expected an error for empty CSV, but got nil")
	}

	if len(questions) != 0 {
		t.Fatalf("Expected 0 questions, but got %d", len(questions))
	}
}

// TestLoadQuestions_MalformedCSV // test that an error is thrown when loading a malformed CSV

// Test Shuffling Functionality
// We can't test randomness itself but can validate that consecutive runs with same questions produce different orders (with seeded randomness if needed)

// TestShuffleQuestions_LengthUnchanged // Test that the length of the questions array remains the same after shuffling
func TestShuffleQuestions_LengthUnchanged(t *testing.T) {
	questions := []Question{
		{"Question 1", "Answer 1"},
		{"Question 2", "Answer 2"},
		{"Question 3", "Answer 3"},
	}

	shuffleQuestions(questions)
	if len(questions) != 3 {
		t.Fatalf("Expected 3 questions after shuffle, got %d", len(questions))
	}
}

// TestShuffleQuestions_OrderChanges // Test that the order of the questions array changes after shuffling but that the CSV file remains the same
func TestShuffleQuestions_OrderChanges(t *testing.T) {
	questions := []Question{
		{"Question 1", "Answer 1"},
		{"Question 2", "Answer 2"},
		{"Question 3", "Answer 3"},
	}

	initialOrder := make([]Question, len(questions))
	copy(initialOrder, questions) // Copy the original order

	shuffleQuestions(questions)

	sameOrder := true
	for i := range questions {
		if questions[i] != initialOrder[i] {
			sameOrder = false
			break
		}
	}

	if sameOrder {
		t.Fatal("Expected shuffled order to be different, but it remained the same")
	}
}

// Test User Input Handling & Answer Checking
// 1. Test exact matches
// 2. Test case-insensitive and whitespace differences
// 3. Test that partial matches are not accepted
// 4. Test that incorrect or non-answers are not accepted
// TestCheckAnswer_ExactMatch // Checks that exact matches are accepted
// TestCheckAnswer_CaseInsensitive // Checks that case-insensitive matches are accepted
// TestCheckAnswer_IgnoreWhitespace // Checks that whitespace is trimmed or ignored
// TestCheckAnswer_WrongAnswer // Checks that incorrect answers are not accepted
// TestCleanInput_TrimSpaces // Checks that input is trimmed of leading/trailing spaces prior to being checked for correctness
// TestCleanInput_ToLowerCase // Checks that input is lowercased prior to being checked for correctness
// TestCleanInput_SpecialCharacters // Checkcases that special characters are removed prior to being checked for correctness

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  Answer  ", "answer"},
		{"ANSWER", "answer"},
		{"AnSwEr", "answer"},
	}

	for _, test := range tests {
		result := cleanInput(test.input)
		if result != test.expected {
			t.Fatalf("Expected %s, but got %s", test.expected, result)
		}
	}
}

// Test Timer Functionality
// `runTimer` function should start a timer and stop it when a channel is closed but is tricky to test. We *can* test its core behavior by checking that elapsed time is updated and that its stopped when signaled.

// TestRunTimer_StartStopCorrectly // Test that the timer starts and stops correctly
func TestTimer_StartStopCorrectly(t *testing.T) {
	timer := time.NewTimer(2 * time.Second)
	stopCh := make(chan bool)

	go func() {
		select {
		case <-timer.C:
			t.Fatal("Timer should have been stopped before it expired")
		case <-stopCh:
			return // Timer stopped as expected
		}
	}()

	time.Sleep(1 * time.Second) // Simulate some processing
	stopCh <- true              // Stop the timer
}

// TestRunTimer_ElapsedTimeCorrect // Test that the elapsed time is correct at different stages of the timer

// Test Quiz Logic
// Test overall quiz logic, ensuring questions are asked correctly, answers are checked, and scores are calculated correctly
// TestQuiz_AllQuestionsAsked // Test that all questions are asked unless time runs out
func TestRunQuiz_AllQuestionsAsked(t *testing.T) {
	questions := []Question{
		{"What is 1 + 1?", "2"},
		{"What is 2 + 2?", "4"},
	}

	correctAnswers := runQuiz(questions, 10*time.Second)
	// Assuming user inputs are simulated elsewhere for production
	if correctAnswers != len(questions) {
		t.Fatalf("Expected %d correct answers, but got %d", len(questions), correctAnswers)
	}
}

// TestQuiz_CorrectAnswerCount // Test that the correct answer count is calculated correctly
// TestQuiz_EndAfterLastQuestion // Test that the quiz ends after the last question is asked or timer runs out
// TestQuiz_EndAfterTimer // Test that the quiz ends after the timer runs out

// Command Line Flags
// Test that we're parsing flags correctly
// TestFlagParsing_ShuffleEnabled // Test that the shuffle flag is parsed correctly
// TestFlagParsing_CustomFilePath // Test that a custom file path is parsed correctly

// Test Error Handling
// Test that errors are handled gracefully like file not found or user input errors
// TestHandleError_FileNotFound // Test that a file not found error is handled correctly
// TestHandleError_InvalidUserInput // Test that invalid user input is handled correctly
