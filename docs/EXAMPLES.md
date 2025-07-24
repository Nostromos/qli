---
sidebar_position: 3
---

# Examples and Use Cases

## Basic Examples

### 1. Simple Math Quiz
Using the default `questions.csv`:
```bash
./quiz
```

Output:
```
Press Enter to start...
Question #1: 5+5
10
Question #2: 7+3
10
...
You scored 11 out of 13 questions correct!
```

### 2. Quick Quiz with Short Timer
Perfect for quick mental math practice:
```bash
./quiz -limit=10s
```

### 3. Randomized Quiz
Different order each time for better learning:
```bash
./quiz -shuffle
```

## Creating Custom Quizzes

### Geography Quiz
Create `geography.csv`:
```csv
What is the capital of Japan?,tokyo
Which country has the most islands?,sweden
What is the smallest country in the world?,vatican city
Which river is the longest in the world?,nile
How many continents are there?,7
```

Run with:
```bash
./quiz -csv=geography.csv -limit=45s
```

### Programming Quiz
Create `programming.csv`:
```csv
What year was Go released?,2009
Who created Linux?,linus torvalds
What does HTTP stand for?,hypertext transfer protocol
What is the time complexity of binary search?,o(log n)
What does SQL stand for?,structured query language
```

Run with extended time:
```bash
./quiz -csv=programming.csv -limit=2m -shuffle
```

### Trivia Quiz with Longer Questions
Create `trivia.csv`:
```csv
"In which year did World War II end?",1945
"What is the largest planet in our solar system?",jupiter
"Who painted the Mona Lisa?",leonardo da vinci
"What is the chemical symbol for gold?",au
"How many strings does a standard guitar have?",6
```

## Advanced Usage Scenarios

### 1. Speed Challenge Mode
Test your speed with a very short timer:
```bash
./quiz -csv=math.csv -limit=15s -shuffle
```

### 2. Study Session Mode
Longer timer for learning new material:
```bash
./quiz -csv=vocabulary.csv -limit=5m
```

### 3. Competition Setup
For classroom or group competitions:
```bash
# Round 1: Easy questions
./quiz -csv=easy.csv -limit=30s

# Round 2: Medium difficulty with shuffling
./quiz -csv=medium.csv -limit=45s -shuffle

# Round 3: Hard questions with more time
./quiz -csv=hard.csv -limit=1m
```

## Tips for Creating Effective Quiz Files

### 1. Answer Formatting
Keep answers simple and consistent:
```csv
What is 2+2?,4
What is two plus two?,4
2 + 2 = ?,4
```

### 2. Handling Special Characters
Questions with commas work correctly:
```csv
"What is 1,000 divided by 10?",100
"In which year did Columbus sail to America?",1492
```

### 3. Case-Insensitive Answers
The quiz accepts various cases:
```csv
What is the capital of France?,paris
```
Valid answers: "Paris", "PARIS", "paris", " paris "

### 4. Multi-Word Answers
Full phrases work well:
```csv
What is the largest mammal?,blue whale
Who wrote Romeo and Juliet?,william shakespeare
```

## Educational Use Cases

### 1. Vocabulary Practice
Create `spanish-vocab.csv`:
```csv
What is 'hello' in Spanish?,hola
What is 'thank you' in Spanish?,gracias
What is 'goodbye' in Spanish?,adios
```

### 2. Math Drills
Create `multiplication.csv`:
```csv
7 x 8,56
9 x 6,54
8 x 7,56
6 x 9,54
```

### 3. Historical Dates
Create `history-dates.csv`:
```csv
When was the Declaration of Independence signed?,1776
When did World War I begin?,1914
When was the internet invented?,1969
```

## Automation and Scripting

### Batch Testing
Create a script to run multiple quizzes:
```bash
#!/bin/bash
echo "Starting Math Quiz..."
./quiz -csv=math.csv -limit=30s
echo "Starting Geography Quiz..."
./quiz -csv=geography.csv -limit=45s
echo "Starting History Quiz..."
./quiz -csv=history.csv -limit=1m
```

### Progress Tracking
Keep a log of scores:
```bash
#!/bin/bash
DATE=$(date +%Y-%m-%d)
echo "Quiz results for $DATE:" >> quiz_log.txt
./quiz -csv=daily.csv | tee -a quiz_log.txt
```

## Common Patterns

### 1. Daily Practice Routine
```bash
# Morning: Quick math warmup
./quiz -csv=math.csv -limit=20s -shuffle

# Afternoon: Subject study
./quiz -csv=science.csv -limit=1m

# Evening: General knowledge
./quiz -csv=trivia.csv -limit=45s -shuffle
```

### 2. Difficulty Progression
Start easy and increase difficulty:
```bash
# Beginner
./quiz -csv=basic.csv -limit=1m

# Intermediate
./quiz -csv=intermediate.csv -limit=45s

# Advanced
./quiz -csv=advanced.csv -limit=30s -shuffle
```

### 3. Subject Rotation
Different subject each day:
```bash
# Monday: Math
./quiz -csv=math.csv

# Tuesday: Science
./quiz -csv=science.csv

# Wednesday: History
./quiz -csv=history.csv

# Thursday: Geography
./quiz -csv=geography.csv

# Friday: Mixed review
./quiz -csv=mixed.csv -shuffle
```