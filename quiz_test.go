package main

import (
	"testing"
	"os"
)

// Quiz File Loading
// 1. Must load questions correctly from a valid CSV file.
// 2. Must properly handle invalid/malformed CSV files.
// 3. Must properly handle empty CSV files.
// TestLoadQuestions_ValidFile // test that questions from a valid CSV are loaded correctly
func TestLoadQuestions_ValidFile(t *testing.T) {
	// test that the path is valid
	t.Run("ValidPath", func (t *testing.T) {
		// create temp file
		tmpFile, err := os.CreateTemp("", "testfiles*.csv")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name()) // clean up after
	})
	// test that the CSV is valid

	// test that the file is loaded correctly

	// test that the questions are loaded correctly

}
// TestLoadQuestions_InvalidFile // test that an error is thrown when loading an invalid path
// TestLoadQuestions_EmptyFile // test that an error is thrown when loading an empty file
// TestLoadQuestions_MalformedCSV // test that an error is thrown when loading a malformed CSV
// TestShuffleQuestions_LengthUnchanged // test that the length of the questions array remains the same after shuffling
// TestShuffleQuestions_OrderChanges // test that the order of the questions array changes after shuffling

// Test Shuffling Functionality
// We can't test randomness itself but can validate that consecutive runs with same questions produce different orders (with seeded randomness if needed)
// TestShuffleQuestions_LengthUnchanged // Test that the length of the questions array remains the same after shuffling
// TestShuffleQuestions_OrderChanges // Test that the order of the questions array changes after shuffling but that the CSV file remains the same


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

// Test Timer Functionality
// `runTimer` function should start a timer and stop it when a channel is closed but is tricky to test. We *can* test its core behavior by checking that elapsed time is updated and that its stopped when signaled. 
// TestRunTimer_StartStopCorrectly // Test that the timer starts and stops correctly
// TestRunTimer_ElapsedTimeCorrect // Test that the elapsed time is correct at different stages of the timer

// Test Quiz Logic
// Test overall quiz logic, ensuring questions are asked correctly, answers are checked, and scores are calculated correctly
// TestQuiz_AllQuestionsAsked // Test that all questions are asked unless time runs out
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
