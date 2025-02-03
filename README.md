# Exercise #1: Quiz Game

## Summary
Build a CLI Quiz app that uses flags to customize the game experience. The quiz is trivia style - the user is asked a question and provides a freeform answer. 

## Thoughts 
**Total Time: 4 hours 14 minutes**
I thought this would be easier. I was so confident in my approach that I spent some time researching CLI formatting and ANSI colors codes. I eventually discarded that approach to better the underlying logic. 

I did learn a lot about golang, packages, and concurrency. Idiomatic go still feels a bit alien, especially with repetitive error handling. Typing `if err != nil {...}` is annoying, but I trust its become idiomatic for a reason. 

https://github.com/Nostromos/qli/blob/6671cb7a4a2bda6751cd9a8775a72fa1bae309f2/quiz.go#L15-L18

I started with multiple structs and was getting really detailed with typing but ended trashing everything but one. 

https://github.com/Nostromos/qli/blob/6671cb7a4a2bda6751cd9a8775a72fa1bae309f2/quiz.go#L115-L135

My initial question/answer loop wouldn't break on timer expiration. In fact it originally had elapsed time saved to a variable that updated each loop. Concurrency is a relatively new concept to me and it took at least an hour for me to realize that `reader.ReadString('\n')` is blocking. It took me another 30-40 mins to figure out that I could use a goroutine to handle the user input portion and a select/case to deal with breaking out of the loop. I'm not sure if this is the best way to do it.

https://github.com/Nostromos/qli/blob/6671cb7a4a2bda6751cd9a8775a72fa1bae309f2/quiz.go#L130

Golang Channels and value assignation are not intuitive to me (coming from the JS/TS world) so this syntax was confusing to learn but I can see how useful it would be. 

## Improvements
- [ ] Add handling for different question types: multiple choice, correct order, etc.
- [ ] Ensure that all loop segments are handled by goroutines. Reading and printing questions are also blocking (I think), so this would make it insignificantly more responsive but probably unnoticeable to the user. It would be good practice. 
- [ ] Remove copilot suggestions for future gophercises. Far too annoying and really hurts ability to learn patterns or reason myself.  

## Requirements
### Part 1
- [x] Ingests CSV file
  - [x] User-provided path via flags
  - [x] Default path and questions included with package (defaults to `problems.csv`)
- [x] Parses CSV into questions and answers
- [x] Asks questions of user and accepts input as answer
- [x] Compares user provided input with answer and keeps track of # of correct/incorrect answers
- [x] Regardless of correct/incorrect, next question is asked immediately after user input
- [x] Quiz Format
  - [x] < 100 questions
  - [x] Single word/number answers
- [x] At end of quiz, program outputs total number of questions correct out of total

### Part 2
- [x] Add timer with default limit of 30 secs, customizable with flag
- [x] User must press `enter` (or another key) before timer starts, then questions are printed out one at a time until an answer is given
- [x] Quiz must stop when time limit is exceeded - do not wait for the user to answer their final question. Ideally stops quiz entirely even if you're waiting on user answer.
- [x] Regardless of whether time limit was exceeded, program must still output "score" with unanswered questions and invalid answers considered incorrect.

### Bonus
- [x] Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. Hint: Check out the strings package.
- [x] Add an option (a new flag) to shuffle the quiz order each time it is run.


<details>
<summary><h2> Exercise details </h2></summary>
This exercise is broken into two parts to help simplify the process of explaining it as well as to make it easier to solve. The second part is harder than the first, so if you get stuck feel free to move on to another problem then come back to part 2 later.

### Part 1
Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

!NOTE: CSV files may have questions with commas in them. Eg: "what 2+2, sir?",4 is a valid row in a CSV. I suggest you look into the CSV package in Go and don’t try to write your own CSV parser.

### Part 2
Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn’t wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

### Bonus
As a bonus exercises you can also…

1. Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. Hint: Check out the strings package.
2. Add an option (a new flag) to shuffle the quiz order each time it is run.
</details>



