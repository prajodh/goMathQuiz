# Quiz Application

This is a simple command-line quiz application written in Go. The quiz questions are loaded from a CSV file, and the user is given a specified amount of time to complete the quiz. The application tracks the number of correct answers and the total number of questions asked.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)

## Features

- Load quiz questions from a CSV file.
- Set a time limit for completing the quiz.
- Non-blocking user input handling.
- Track the number of correct answers.
- Display results at the end of the quiz.

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/goMathQuiz.git
    cd goMathQuiz
    ```

2. **Build the application:**

    ```bash
    go build
    ```

## Usage

Run the application with the following command:

```bash
./quiz-app -csv=path/to/quiz.csv -time=30
