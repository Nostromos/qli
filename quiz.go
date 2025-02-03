package main

import (
	"fmt"
  "encoding/csv"
	"bufio"
	"os"
	"strings"
	"flag"
	"time"
	"Math/rand"
)

func main () {
	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("You'll be asked a series of questions and you'll need to answer them before the time limit (30 secs by default)")
	fmt.Println("Quizzes will be less than 100 questions and you'll have single word/number answers")
	fmt.Println("At the end, you'll see how many you got right. Good luck!")

	fmt.Println("Please enter path to quiz questions or hit enter to use default questions...")

  shuffle := flag.Bool("shuffle", false, "Shuffle the quiz questions")
	var defaultQuestions string = "./questions.csv";
	var questionsPath = bufio.NewReader(os.Stdin)
	path, err := questionsPath.ReadString('\n')
	if err != nil {
		panic(err)
	} else if path == "\n" {
		path = defaultQuestions
	}

	path = strings.TrimSpace(path)
	
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	if *shuffle {
		data = shuffleQuestions(data)
	}
	var correctAnswers int = 0
	var totalQuestions int = len(data)

	reader2 := bufio.NewReader(os.Stdin)

	fmt.Println("Press Enter to start the quiz...")
	reader2.ReadString('\n') // Wait for user to press Enter

	// Record start time for both timers
	startTimer1 := time.Now()

	// Create a channel to stop the timers
	stop := make(chan bool)

	// Launch goroutines for both timers
	go runTimer(startTimer1, 1, stop, Red)   // Timer 1 in red on line 1

	fmt.Println("Press Enter to stop the timers...")
	reader2.ReadString('\n') // Wait for Enter to stop

	// Signal both goroutines to stop
	stop <- true
	stop <- true

	for _, question := range data {
		var correct = questionsAnswer(question[0], question[1])
		if correct {
			correctAnswers++
		}
	}

	fmt.Printf("You got %d out of %d questions correct!\n", correctAnswers, totalQuestions)
}

func questionsAnswer (q, a string) bool {
	fmt.Println(q)
	var answer string
	fmt.Scanln(&answer)
	answer = cleanInput(answer)
	a = cleanInput(a)
	if answer == a {
		return true
	} else {
		return false;
	}
}

func cleanInput (s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func shuffleQuestions (data [][]string) [][]string {
	for i := range data {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func runTimer(start time.Time, line int, stop chan bool, color string) {
	for {
		select {
		case <-stop:
			return // Stop the timer when signaled
		default:
			elapsed := time.Since(start)
			// Use ANSI escape code to move cursor and update the correct line with color
			fmt.Printf("\033[%d;0H%sElapsed time for Timer %d: %s%s", line, color, line, elapsed.Round(time.Second), Reset)
			time.Sleep(100 * time.Millisecond) // Update every 100ms
		}
	}
}
